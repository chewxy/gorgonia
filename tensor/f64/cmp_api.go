package tensorf64

// public API for comparison ops

import "github.com/chewxy/gorgonia/tensor/types"

// Lt performs a pointwise less than comparison (a < b). a and b can either be float64 or *Tensor.
// It returns a *tensorbool.Tensor, NOT a *Tensor. This is important
//
// If both operands are *Tensor, shape is checked first.
// Even though the underlying data may have the same size (say (2,2) vs (4,1)), if they have different shapes, it will error out.
func Lt(a, b interface{}, opts ...types.FuncOpt) (retVal types.Tensor, err error) {
	boolT := !parseAsFloat64(opts...)

	at, atok := a.(*Tensor)
	bt, btok := b.(*Tensor)
	af, afok := a.(float64)
	bf, bfok := b.(float64)
	op := lt

	switch {
	case atok && btok:
		return at.tensorCmp(op, bt, boolT)
	case boolT && atok && bfok:
		return at.scalarCmp(op, true, bf)
	case boolT && afok && btok:
		return bt.scalarCmp(op, false, af)

	// returns TensorF64
	case !boolT && atok && bfok:
		var b []bool
		if b, err = scalarCmpBacking(op, true, bf, at.data); err == nil {
			backing := boolsToFloat64s(b)
			retVal = NewTensor(WithShape(at.Shape()...), WithBacking(backing))
		}
		return
	case !boolT && afok && btok:
		var b []bool
		if b, err = scalarCmpBacking(op, false, af, bt.data); err == nil {
			backing := boolsToFloat64s(b)
			retVal = NewTensor(WithShape(bt.Shape()...), WithBacking(backing))
		}
		return
	default:
		err = types.NewError(types.DtypeMismatch, "Comparison cannot be done on %T and %T", a, b)
		return
	}
	panic("unreachable")
}

// Gt performs a pointwise greater than comparison (a > b). a and b can either be float64 or *Tensor.
// It returns a *tensorbool.Tensor, NOT a *Tensor. This is important
//
// If both operands are *Tensor, shape is checked first.
// Even though the underlying data may have the same size (say (2,2) vs (4,1)), if they have different shapes, it will error out.
func Gt(a, b interface{}, opts ...types.FuncOpt) (retVal types.Tensor, err error) {
	boolT := !parseAsFloat64(opts...)

	at, atok := a.(*Tensor)
	bt, btok := b.(*Tensor)
	af, afok := a.(float64)
	bf, bfok := b.(float64)
	op := gt

	switch {
	case atok && btok:
		return at.tensorCmp(op, bt, boolT)
	case boolT && atok && bfok:
		return at.scalarCmp(op, true, bf)
	case boolT && afok && btok:
		return bt.scalarCmp(op, false, af)
	case !boolT && atok && bfok:
		var b []bool
		if b, err = scalarCmpBacking(op, true, bf, at.data); err == nil {
			backing := boolsToFloat64s(b)
			retVal = NewTensor(WithShape(at.Shape()...), WithBacking(backing))
		}
		return
	case !boolT && afok && btok:
		var b []bool
		if b, err = scalarCmpBacking(op, false, af, bt.data); err == nil {
			backing := boolsToFloat64s(b)
			retVal = NewTensor(WithShape(bt.Shape()...), WithBacking(backing))
		}
		return
	default:
		err = types.NewError(types.DtypeMismatch, "Comparison cannot be done on %T and %T. %v, atok  %vbtok %v", a, b, boolT, atok, btok)
		return
	}
	panic("unreachable")
}

// Lte performs a pointwise less than eq comparison (a <= b). a and b can either be float64 or *Tensor.
// It returns a *tensorbool.Tensor, NOT a *Tensor. This is important
//
// If both operands are *Tensor, shape is checked first.
// Even though the underlying data may have the same size (say (2,2) vs (4,1)), if they have different shapes, it will error out.
func Lte(a, b interface{}, opts ...types.FuncOpt) (retVal types.Tensor, err error) {
	boolT := !parseAsFloat64(opts...)

	at, atok := a.(*Tensor)
	bt, btok := b.(*Tensor)
	af, afok := a.(float64)
	bf, bfok := b.(float64)
	op := lte

	switch {
	case atok && btok:
		return at.tensorCmp(op, bt, boolT)
	case boolT && atok && bfok:
		return at.scalarCmp(op, true, bf)
	case boolT && afok && btok:
		return bt.scalarCmp(op, false, af)

	// returns TensorF64
	case !boolT && atok && bfok:
		var b []bool
		if b, err = scalarCmpBacking(op, true, bf, at.data); err == nil {
			backing := boolsToFloat64s(b)
			retVal = NewTensor(WithShape(at.Shape()...), WithBacking(backing))
		}
		return
	case !boolT && afok && btok:
		var b []bool
		if b, err = scalarCmpBacking(op, false, af, bt.data); err == nil {
			backing := boolsToFloat64s(b)
			retVal = NewTensor(WithShape(bt.Shape()...), WithBacking(backing))
		}
		return
	default:
		err = types.NewError(types.DtypeMismatch, "Comparison cannot be done on %T and %T", a, b)
		return
	}
	panic("unreachable")
}

// Gte performs a pointwise greater than eq comparison (a >= b). a and b can either be float64 or *Tensor.
// It returns a *tensorbool.Tensor, NOT a *Tensor. This is important
//
// If both operands are *Tensor, shape is checked first.
// Even though the underlying data may have the same size (say (2,2) vs (4,1)), if they have different shapes, it will error out.
func Gte(a, b interface{}, opts ...types.FuncOpt) (retVal types.Tensor, err error) {
	boolT := !parseAsFloat64(opts...)

	at, atok := a.(*Tensor)
	bt, btok := b.(*Tensor)
	af, afok := a.(float64)
	bf, bfok := b.(float64)
	op := gte

	switch {
	case atok && btok:
		return at.tensorCmp(op, bt, boolT)
	case boolT && atok && bfok:
		return at.scalarCmp(op, true, bf)
	case boolT && afok && btok:
		return bt.scalarCmp(op, false, af)

	// returns TensorF64
	case !boolT && atok && bfok:
		var b []bool
		if b, err = scalarCmpBacking(op, true, bf, at.data); err == nil {
			backing := boolsToFloat64s(b)
			retVal = NewTensor(WithShape(at.Shape()...), WithBacking(backing))
		}
		return
	case !boolT && afok && btok:
		var b []bool
		if b, err = scalarCmpBacking(op, false, af, bt.data); err == nil {
			backing := boolsToFloat64s(b)
			retVal = NewTensor(WithShape(bt.Shape()...), WithBacking(backing))
		}
		return
	default:
		err = types.NewError(types.DtypeMismatch, "Comparison cannot be done on %T and %T", a, b)
		return
	}
	panic("unreachable")
}

// Eq performs a pointwise equality comparison (a == b). a and b can either be float64 or *Tensor.
// It returns a *tensorbool.Tensor, NOT a *Tensor. This is important
//
// If both operands are *Tensor, shape is checked first.
// Even though the underlying data may have the same size (say (2,2) vs (4,1)), if they have different shapes, it will error out.
func Eq(a, b interface{}, opts ...types.FuncOpt) (retVal types.Tensor, err error) {
	boolT := !parseAsFloat64(opts...)

	at, atok := a.(*Tensor)
	bt, btok := b.(*Tensor)
	af, afok := a.(float64)
	bf, bfok := b.(float64)
	op := eq

	switch {
	case atok && btok:
		return at.tensorCmp(op, bt, boolT)
	case boolT && atok && bfok:
		return at.scalarCmp(op, true, bf)
	case boolT && afok && btok:
		return bt.scalarCmp(op, false, af)

	// returns TensorF64
	case !boolT && atok && bfok:
		var b []bool
		if b, err = scalarCmpBacking(op, true, bf, at.data); err == nil {
			backing := boolsToFloat64s(b)
			retVal = NewTensor(WithShape(at.Shape()...), WithBacking(backing))
		}
		return
	case !boolT && afok && btok:
		var b []bool
		if b, err = scalarCmpBacking(op, false, af, bt.data); err == nil {
			backing := boolsToFloat64s(b)
			retVal = NewTensor(WithShape(bt.Shape()...), WithBacking(backing))
		}
		return
	default:
		err = types.NewError(types.DtypeMismatch, "Comparison cannot be done on %T and %T", a, b)
		return
	}
	panic("unreachable")
}

// Ne performs a pointwise equality comparison (a != b). a and b can either be float64 or *Tensor.
// It returns a *tensorbool.Tensor, NOT a *Tensor. This is important
//
// If both operands are *Tensor, shape is checked first.
// Even though the underlying data may have the same size (say (2,2) vs (4,1)), if they have different shapes, it will error out.
func Ne(a, b interface{}, opts ...types.FuncOpt) (retVal types.Tensor, err error) {
	boolT := !parseAsFloat64(opts...)

	at, atok := a.(*Tensor)
	bt, btok := b.(*Tensor)
	af, afok := a.(float64)
	bf, bfok := b.(float64)
	op := ne

	switch {
	case atok && btok:
		return at.tensorCmp(op, bt, boolT)
	case boolT && atok && bfok:
		return at.scalarCmp(op, true, bf)
	case boolT && afok && btok:
		return bt.scalarCmp(op, false, af)

	// returns TensorF64
	case !boolT && atok && bfok:
		var b []bool
		if b, err = scalarCmpBacking(op, true, bf, at.data); err == nil {
			backing := boolsToFloat64s(b)
			retVal = NewTensor(WithShape(at.Shape()...), WithBacking(backing))
		}
		return
	case !boolT && afok && btok:
		var b []bool
		if b, err = scalarCmpBacking(op, false, af, bt.data); err == nil {
			backing := boolsToFloat64s(b)
			retVal = NewTensor(WithShape(bt.Shape()...), WithBacking(backing))
		}
		return
	default:
		err = types.NewError(types.DtypeMismatch, "Comparison cannot be done on %T and %T", a, b)
		return
	}
	panic("unreachable")
}
