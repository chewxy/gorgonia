package tensor

import "sync"

var habbo sync.Mutex
var usePool bool = true

// tensorPool is a pool of *Tensor grouped by size. It's guarded by poolsClosed
var poolsClosed sync.RWMutex
var densePool map[dtype]map[int]*sync.Pool = make(map[dtype]map[int]*sync.Pool)

const (
	maxAPDims = 8
)

// UsePool enables the use of a pool of *Tensors as provided in the package. This is the default option
func UsePool() {
	habbo.Lock()
	usePool = true
	habbo.Unlock()
}

// DontUsePool makes sure the functions don't use the tensor pool provided.
// This is useful as certain applications don't lend themselves well to use of the pool.
// Examples of such applications would be one where many tensors of wildly different sizes are created all the time.
func DontUsePool() {
	habbo.Lock()
	usePool = false
	habbo.Unlock()
}

func newDensePool(dt dtype, size int) *sync.Pool {
	var pool *sync.Pool

	poolsClosed.Lock()
	// check once more that before the lock was acquired, that nothing else had written to that key
	pools, ok := densePool[dt]
	if !ok {
		pools = make(map[int]*sync.Pool)
		densePool[dt] = pools
	}

	if p, ok := pools[size]; !ok {
		pool = new(sync.Pool)
		l := size
		t := dt

		pool.New = func() interface{} {
			return newDense(t, l)
		}

		pools[size] = pool
	} else {
		pool = p
	}

	poolsClosed.Unlock()
	return pool
}

func borrowDense(dt dtype, size int) *Dense {
	if !usePool {
		return newDense(dt, size)
	}

	var pool *sync.Pool

	poolsClosed.RLock()
	pools, ok := densePool[dt]
	poolsClosed.RUnlock()

	if !ok {
		pool = newDensePool(dt, size)
		goto end
	}

	if pool, ok = pools[size]; !ok {
		pool = newDensePool(dt, size)
	}

end:
	return pool.Get().(*Dense)
}

// ReturnTensor returns a Tensor to their respective pools. Use with caution
func ReturnTensor(t Tensor) {
	if !usePool {
		return
	}
	switch tt := t.(type) {
	case *Dense:
		dt, ok := tt.t.(dtype)
		if !ok {
			return
		}
		if _, ok := densePool[dt]; !ok {
			return
		}

		if tt.viewOf != nil {
			ReturnAP(tt.AP)
			tt.AP = nil
			if tt.old != nil {
				ReturnAP(tt.old)
				tt.old = nil
			}
			if tt.transposeWith != nil {
				ReturnInts(tt.transposeWith)
				tt.transposeWith = nil
			}
			tt.data = nil
			return // yes, we're not putting it back into the pool
		}

		size := tt.data.Cap()
		poolsClosed.RLock()
		pool, ok := densePool[dt][size]
		poolsClosed.RUnlock()
		if !ok {
			pool = newDensePool(dt, size)
		}

		if tt.old != nil {
			ReturnAP(tt.old)
			tt.old = nil
		}

		if tt.transposeWith != nil {
			ReturnInts(tt.transposeWith)
			tt.transposeWith = nil
		}

		tt.Unlock()
		pool.Put(tt)
	}
}

/* AP POOL */

// apPool supports tensors up to 4-dimensions. Because, c'mon, you're not likely to use anything more than 5
var apPool [maxAPDims]sync.Pool

func BorrowAP(dims int) *AP {
	if dims >= maxAPDims {
		ap := new(AP)
		ap.shape = make(Shape, dims)
		ap.strides = make([]int, dims)
		return ap
	}

	ap := apPool[dims].Get().(*AP)

	// restore strides and shape to whatever that may have been truncated
	ap.strides = ap.strides[:cap(ap.strides)]
	return ap
}

func ReturnAP(ap *AP) {
	if ap.Dims() >= maxAPDims {
		return
	}
	apPool[ap.Dims()].Put(ap)
}

/* INTS POOL */

var intsPool [8]sync.Pool

func init() {
	for i := range intsPool {
		size := i
		intsPool[i].New = func() interface{} { return make([]int, size) }
	}

	for i := range apPool {
		l := i
		apPool[i].New = func() interface{} {
			ap := new(AP)
			ap.strides = make([]int, l)
			ap.shape = make(Shape, l)
			return ap
		}
	}
}

func BorrowInts(size int) []int {
	if size >= 8 {
		return make([]int, size)
	}

	retVal := intsPool[size].Get()
	if retVal == nil {
		return make([]int, size)
	}
	return retVal.([]int)
}

func ReturnInts(is []int) {
	if is == nil {
		return
	}
	size := cap(is)
	if size >= 8 {
		return
	}
	is = is[:cap(is)]
	for i := range is {
		is[i] = 0
	}

	intsPool[size].Put(is)
}
