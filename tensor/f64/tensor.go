package tensorf64

import (
	"fmt"

	"github.com/chewxy/gorgonia/tensor/types"
)

type Tensor struct {
	*types.AP
	data []float64

	// backup AP. When a transpose is done, the old *AP is backed up here, for easy untransposes
	old           *types.AP
	transposeWith []int

	// if viewOf != nil, then this *Tensor is a view.
	viewOf *Tensor
}

// a ConsOpt is a tensor construction option
type consOpt func(*Tensor)

func (c consOpt) Opt() {}

// NewTensor creates a new Float64 *Tensor
func NewTensor(opts ...types.ConsOpt) *Tensor {
	t := new(Tensor)
	t.AP = new(types.AP)

	for _, opt := range opts {
		o := opt.(consOpt)
		o(t)
	}
	t.fix()
	// TODO: sanity check
	if err := t.sanity(); err != nil {
		panic(err)
	}
	return t
}

// newBorrowedTensor tries to borrow from the tensor pool. It isn't zeroed!
func newBorrowedTensor(size int, opts ...consOpt) *Tensor {
	t := BorrowTensor(size)

	for _, opt := range opts {
		opt(t)
	}

	t.fix()
	if err := t.sanity(); err != nil {
		panic(err)
	}
	return t
}

func newTensor(size int) *Tensor {
	t := new(Tensor)
	t.AP = new(types.AP)
	t.setShape(size)
	t.data = make([]float64, size)
	return t
}

// Ones create a ndarray of the given shape, and fills it with 1.0
func Ones(shape ...int) *Tensor {
	if len(shape) == 0 {
		one := float64(1) //@DEFAULTONE
		return NewTensor(AsScalar(one))
	}

	t := BorrowTensor(types.Shape(shape).TotalSize())
	for i := range t.data {
		t.data[i] = float64(1) //@DEFAULTONE
	}

	t.setShape(shape...)
	return t
}

// Zeroes create a ndarray of a given shape and fills it with float64(0) (which is Go's default value)
// It's here mainly as a convenience function
func Zeroes(shape ...int) *Tensor {
	t := BorrowTensor(types.Shape(shape).TotalSize())
	t.setShape(shape...)
	t.Zero()
	return t
}

// I creates the identity matrix (usually a square) matrix with 1s across the diagonals, and zeroes elsewhere, like so:
//		Matrix(4,4)
// 		⎡1  0  0  0⎤
// 		⎢0  1  0  0⎥
// 		⎢0  0  1  0⎥
// 		⎣0  0  0  1⎦
// While technically an identity matrix is a square matrix, in attempt to keep feature parity with Numpy,
// the I() function allows you to create non square matrices, as well as an index to start the diagonals.
//
// For example:
//		T = I(4, 4, 1)
// Yields:
//		⎡0  1  0  0⎤
//		⎢0  0  1  0⎥
//		⎢0  0  0  1⎥
//		⎣0  0  0  0⎦
//
// The index k can also be a negative number:
// 		T = I(4, 4, -1)
// Yields:
// 		⎡0  0  0  0⎤
// 		⎢1  0  0  0⎥
// 		⎢0  1  0  0⎥
// 		⎣0  0  1  0⎦
func I(r, c, k int) (retVal *Tensor) {
	retVal = NewTensor(WithShape(r, c))
	if k >= c {
		return
	}

	i := k
	if k < 0 {
		i = (-k) * c
	}

	var s *Tensor
	var err error
	end := c - k
	if end > r {
		s, err = retVal.Slice(nil)
	} else {
		s, err = retVal.Slice(rs{0, end, 1})
	}
	defer ReturnTensor(s)

	if err != nil {
		panic(err)
	}

	var nexts []int
	iter := types.NewFlatIterator(s.AP)
	nexts, err = iter.Slice(rs{i, s.Size(), c + 1})

	for _, v := range nexts {
		s.data[v] = float64(1) //@DEFAULTONE
	}
	return
}

// WithBacking is a construction option for NewTensor
// Use it as such:
//		backing := []float64{1,2,3,4}
// 		t := NewTensor(WithBacking(backing))
// It can be used with other construction options like WithShape
func WithBacking(a []float64) consOpt {
	f := func(t *Tensor) {
		t.data = a
	}
	return f
}

// WithShape is a construction option for NewNDArray - it creates the ndarray in the required shape
func WithShape(dims ...int) consOpt {
	f := func(t *Tensor) {
		throw := types.BorrowInts(len(dims))
		copy(throw, dims)
		t.setShape(throw...)

		// special case for scalars
		if len(throw) == 0 {
			t.data = make([]float64, 1)
		}
	}

	return consOpt(f)
}

// AsScalar is a construction option for representing a scalar value as an ndarray
func AsScalar(s float64) consOpt {
	f := func(t *Tensor) {
		t.setShape()
		t.data = []float64{s}
	}
	return f
}

func (t *Tensor) setShape(s ...int) {
	t.Unlock()
	t.SetShape(s...)
	t.Lock()
	return
}

func (t *Tensor) fix() {
	if t.Shape() == nil {
		if t.data == nil {
			return
		}
		// otherwise, the shape is automatically a [n,1]
		rows := len(t.data)
		if rows == 1 {
			t.SetShape() // it's a scalar!
		} else {
			t.SetShape(rows) // it's a vector (unknown whether column or row)
		}
	}

	if t.data == nil {
		size := t.Shape().TotalSize()
		t.data = make([]float64, size)
	}
	t.Lock() // don't put this in a defer - if t.data == nil and t.Shape() == nil. then leave it unlocked
}

// sanity is a function that sanity checks that a tensor is correct.
func (t *Tensor) sanity() error {
	if t.AP != nil && t.Shape() == nil && t.data == nil {
		return types.EmptyTensorError()
	}

	size := len(t.data)
	expected := t.Size()
	if t.viewOf == nil && size != expected && !t.IsScalar() {
		return types.NewError(types.ShapeMismatch, "Expected backing data to have %d elements from shape %v. Got %d instead", expected, t.Shape(), size)
	}
	// TODO: sanity check for views
	return nil
}

func (t *Tensor) oshape() types.Shape {
	if t.old != nil {
		return t.old.Shape()
	}
	return t.Shape()
}

func (t *Tensor) ostrides() []int {
	if t.old != nil {
		return t.old.Strides()
	}
	return t.Strides()
}

func (t *Tensor) Info() *types.AP    { return t.AP }
func (t *Tensor) Dtype() types.Dtype { return types.Float64 }
func (t *Tensor) DataSize() int      { return len(t.data) }

// Reshape reshapes a *Tensor. If the tensors need to be materialized (either it's a view or transpose), it will be materialized before the reshape happens
func (t *Tensor) Reshape(dims ...int) error {
	if t.viewOf != nil {
		return notyetimplemented("Reshape", "views")
	}

	if t.old != nil {
		t.Transpose()
	}

	return t.reshape(dims...)
}

func (t *Tensor) reshape(dims ...int) error {
	t.setShape(dims...)
	return t.sanity()
}

// Zero zeroes a *Tensor.
func (t *Tensor) Zero() {
	for i := range t.data {
		t.data[i] = float64(0) //@DEFAULTZERO
	}
}

func (t *Tensor) SetAll(val interface{}) error {
	if val == 1 {
		for i := range t.data {
			t.data[i] = float64(1) //@DEFAULTONE
		}
		return nil
	}

	v, ok := val.(float64)
	if !ok {
		return types.NewError(types.DtypeMismatch, "Cannot set val of %T. Expected Float64", val)
	}

	for i := range t.data {
		t.data[i] = v
	}
	return nil
}

// ScalarValue() returns the scalar value of a *Tensor,
// IF and ONLY IF it's a Tensor representation of a scalar value.
// This is required because operations like a (vec · vec) would return a scalar value.
// I didn't want to return interface{} for all the API methods, so the next best solution is to
// wrap the scalar value in a *Tensor
func (t *Tensor) ScalarValue() interface{} {
	if !t.IsScalar() {
		panic(fmt.Sprintf("ScalarValue only works when the Tensor is a representation of a scalar value. The value of the tensor is %v", t))
	}

	return t.data[0]
}

// Eq checks that two types.Tensor are equal. If the shapes are the same, but the strides are not the same, it's will still be considered the same
func (t *Tensor) Eq(other types.Tensor) bool {
	if ot, ok := other.(*Tensor); ok {
		if ot == t {
			return true
		}

		if len(ot.data) != len(t.data) {
			return false
		}

		for i, v := range t.data {
			if ot.data[i] != v {
				return false
			}
		}

		if !t.Shape().Eq(ot.Shape()) {
			return false
		}
		//TODO: MORE METADATA CHECKS!

		return true
	}
	return false
}

// Clone clones the *Tensor. It creates a copy of the data. A new underlying array will be allocated
func (t *Tensor) Clone() *Tensor {
	retVal := new(Tensor)
	retVal.AP = t.AP.Clone()
	if t.old != nil {
		retVal.old = t.old.Clone()
	}

	if t.transposeWith != nil {
		retVal.transposeWith = types.BorrowInts(len(t.transposeWith))
		for i, v := range t.transposeWith {
			retVal.transposeWith[i] = v
		}
	}

	var newData []float64
	if t.viewOf != nil {
		iter := types.NewFlatIterator(t.AP)
		newData = make([]float64, t.Shape().TotalSize())
		newData = newData[:0]
		for i, err := iter.Next(); err == nil; i, err = iter.Next() {
			newData = append(newData, t.data[i])
		}

	} else {
		newData = make([]float64, len(t.data))
		copy(newData, t.data)
	}

	retVal.data = newData
	// retVal.viewOf = t
	// retVal.Lock()
	return retVal
}

func (t *Tensor) borrowClone() *Tensor {
	retVal := BorrowTensor(len(t.data))
	types.ReturnAP(retVal.AP)
	retVal.AP = t.AP.Clone()

	if t.old != nil {
		retVal.old = t.old.Clone()
	}

	newdata := make([]float64, len(t.data))
	copy(newdata, t.data)
	retVal.data = newdata
	retVal.Lock()
	return retVal
}

//  IsView indicates if the Tensor is a view of another (typically from slicing)
func (t *Tensor) IsView() bool {
	return t.viewOf != nil
}

// IsMaterializeable() indicates if the Tensor is materializable - if it has either gone through some transforms or slicing
func (t *Tensor) IsMaterializable() bool {
	return t.viewOf != nil || t.old != nil
}

/* Misc public API */
func (t *Tensor) Data() interface{} { return t.data }

func assignArray(dest, src *Tensor) (err error) {
	// var copiedSrc bool

	if src.IsScalar() {
		panic("HELP")
	}

	dd := dest.Dims()
	sd := src.Dims()

	ds := dest.Strides()
	ss := src.Strides()

	// when dd == 1, and the strides point in the same direction
	// we copy to a temporary if there is an overlap of data
	if ((dd == 1 && sd >= 1 && ds[0]*ss[sd-1] < 0) || dd > 1) && overlaps(dest.data, src.data) {
		// create temp
		// copiedSrc = true
	}

	// broadcast src to dest for raw iteration
	tmpShape := types.Shape(types.BorrowInts(sd))
	tmpStrides := types.BorrowInts(len(src.Strides()))
	copy(tmpShape, src.Shape())
	copy(tmpStrides, src.Strides())
	defer types.ReturnInts(tmpShape)
	defer types.ReturnInts(tmpStrides)

	if sd > dd {
		tmpDim := sd
		for tmpDim > dd && tmpShape[0] == 1 {
			tmpDim--

			// this is better than tmpShape = tmpShape[1:]
			// because we are going to return these ints later
			copy(tmpShape, tmpShape[1:])
			copy(tmpStrides, tmpStrides[1:])
		}
	}

	var newStrides []int
	if newStrides, err = types.BroadcastStrides(dest.Shape(), tmpShape, ds, tmpStrides); err != nil {
		return
	}

	dap := dest.AP
	sap := types.NewAP(tmpShape, newStrides)

	diter := types.NewFlatIterator(dap)
	siter := types.NewFlatIterator(sap)
	// dch := diter.Chan()
	// sch := siter.Chan()

	var i, j int
	// var ok bool
	for {
		// if i, ok = <-dch; !ok {
		// 	break
		// }
		// if j, ok = <-sch; !ok {
		// 	break
		// }

		if i, err = diter.Next(); err != nil {
			if _, ok := err.(NoOpError); !ok {
				return err
			}
			err = nil
			break
		}
		if j, err = siter.Next(); err != nil {
			if _, ok := err.(NoOpError); !ok {
				return err
			}
			err = nil
			break
		}

		dest.data[i] = src.data[j]
	}

	return
}

// courtesy of roger pepe: https://groups.google.com/d/msg/golang-nuts/_Pj9S_Ljp9o/GMo5uPzHbeAJ
func overlaps(a, b []float64) bool {
	if cap(a) == 0 || cap(b) == 0 {
		return false
	}
	if &a[0:cap(a)][cap(a)-1] != &b[0:cap(b)][cap(b)-1] {
		return false
	}

	a0 := -cap(a)
	a1 := a0 + len(a)
	b0 := -cap(b)
	b1 := b0 + len(b)
	return a1 > b0 && b1 > a0
}

/* Other Data types */

// rs is a struct representing a ranged slice: [start:end:step]
type rs struct {
	start, end, step int
}

// makeRS creates a ranged slice. It takes an optional step param.
func makeRS(start, end int, opts ...int) rs {
	step := 1
	if len(opts) > 0 {
		step = opts[0]
	}
	return rs{
		start: start,
		end:   end,
		step:  step,
	}
}

func (s rs) Start() int { return s.start }
func (s rs) End() int   { return s.end }
func (s rs) Step() int  { return s.step }

// ss is a single slice, representing this: [start:start+1:0]
type ss int

func (s ss) Start() int { return int(s) }
func (s ss) End() int   { return int(s) + 1 }
func (s ss) Step() int  { return 0 }
