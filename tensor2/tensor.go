package tensor

import (
	"encoding/gob"
	"fmt"
	"io"

	"github.com/pkg/errors"
)

type Tensor interface {
	// info about the ndarray
	Info() *AP
	Shape() Shape
	Strides() []int
	Dtype() Dtype
	Dims() int
	Size() int
	DataSize() int

	// ops
	At(...int) (interface{}, error)
	Reshape(...int) error
	T(axes ...int) error
	UT()
	Transpose() // Transpose actually moves the data

	// data related interface
	Zeroer
	MemSetter
	Dataer
	Eq
	Cloner

	// type overloading shit
	IsScalar() bool
	ScalarValue() interface{}

	// view related shit
	IsView() bool
	Materialize() Tensor

	// formatters
	fmt.Formatter
	fmt.Stringer

	// all Tensors are serializable to these formats
	WriteNpy(io.Writer) error
	ReadNpy(io.Reader) error
	gob.GobEncoder
	gob.GobDecoder
}

// Dotter is used to implement sparse matrices
type Dotter interface {
	Tensor
	Dot(Tensor, ...FuncOpt) (Tensor, error)
}

// New creates a new Dense Tensor. For sparse arrays use their relevant construction function
func New(opts ...ConsOpt) *Dense {
	d := new(Dense)
	d.AP = new(AP)
	for _, opt := range opts {
		opt(d)
	}
	d.fix()
	if err := d.sanity(); err != nil {
		panic(err)
	}

	return d
}

// Ones create a ndarray of the given shape, and fills it with 1.0
func Ones(dt Dtype, shape ...int) Tensor {
	if len(shape) == 0 {
		d := newDense(dt, Shape(shape).TotalSize())
		d.setShape() // scalar shape
		if o, ok := d.data.(Oner); ok {
			o.One()
		} else {
			// TODO
		}
		return d
	}

	if t, ok := dt.(dtype); ok {
		d := borrowDense(t, Shape(shape).TotalSize())
		if o, ok := d.data.(Oner); ok {
			o.One()
			d.setShape(shape...)
			return d
		}
		panic("TODO")
	}
	panic("Unreachable")
}

// Zeroes create a ndarray of a given shape and fills it with float64(0) (which is Go's default value)
// It's here mainly as a convenience function
// func Zeroes(dt Dtype, shape ...int) Tensor {

// 	d := newDense(dt, Shape(shape).TotalSize())
// 	d.setShape(shape...)
// 	d.Zero()
// 	return d
// }

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

// func I(dt Dtype, r, c, k int) (retVal Tensor) {
// 	d := borrowDense(dt, r*c)
// 	d.reshape(r, c)

// 	if k >= c {
// 		return
// 	}

// 	i := k
// 	if k < 0 {
// 		i = (-k) * c
// 	}

// 	var s *Dense
// 	var err error
// 	end := c - k
// 	if end > r {
// 		s, err = d.Slice(nil)
// 	} else {
// 		s, err = d.Slice(rs{0, end, 1})
// 	}
// 	defer ReturnTensor(s)

// 	if err != nil {
// 		panic(err)
// 	}

// 	var nexts []int
// 	iter := NewFlatIterator(s.AP)
// 	nexts, err = iter.Slice(rs{i, s.Size(), c + 1})

// 	for _, v := range nexts {

// 		s.data[v] = float64(1) //@DEFAULTONE
// 	}
// 	return d
// }

func getDense(t Tensor) (*Dense, error) {
	if t == nil {
		return nil, nil
	}

	if retVal, ok := t.(*Dense); ok {
		return retVal, nil
	}
	return nil, errors.Errorf(extractionFail, "*Dense", t)
}

// getFloatDense extracts a *Dense from a Tensor and ensures that the .data is a Array that implements Float
func getFloatDense(t Tensor) (retVal *Dense, err error) {
	if t == nil {
		return
	}

	if retVal, err = getDense(t); err != nil {
		err = errors.Wrapf(err, opFail, "getFloatDense")
		return
	}
	if f, ok := retVal.data.(Float); !ok {
		err = errors.Errorf(dtypeMismatch, f, retVal.data)
		return
	}
	return
}
