package gorgonia

import (
	"fmt"

	"github.com/pkg/errors"
)

// BinaryXent is a convenience function for doing binary crossentropy stuff.
// The formula is as below:
// 		-(y * logprob) +  (1-y)(1-logprob)
func BinaryXent(output, target *Node) (retVal *Node, err error) {
	var one *Node
	var logO, omt, omo, tLogO *Node

	// which constant one to use?
	var dt Dtype
	if dt, err = dtypeOf(output.t); err != nil {
		return nil, errors.Wrapf(err, dtypeExtractionFail, output.t)
	}

	switch dt {
	case Float64:
		one = onef64
	case Float32:
		one = onef32
	default:
		return nil, errors.Errorf(nyiFail, "BinaryXEnt", dt)
	}

	if logO, err = Log(output); err != nil {
		return nil, errors.Wrap(err, operationError)
	}

	if omt, err = Sub(one, target); err != nil {
		return nil, errors.Wrap(err, operationError)
	}

	if omo, err = Sub(one, output); err != nil {
		return nil, errors.Wrap(err, operationError)
	}

	if tLogO, err = HadamardProd(target, logO); err != nil {
		return nil, errors.Wrap(err, operationError)
	}

	if retVal, err = Log(omo); err != nil {
		return nil, errors.Wrap(err, operationError)
	}

	if retVal, err = HadamardProd(omt, retVal); err != nil {
		return nil, errors.Wrap(err, operationError)
	}

	if retVal, err = Add(tLogO, retVal); err != nil {
		return nil, errors.Wrap(err, operationError)
	}

	return Neg(retVal)
}

// Dropout is a convenience function to implement dropout.
// It uses randomly zeroes out a *Tensor with a probabilty drawn from
// a uniform distribution
func Dropout(x *Node, prob float64) (retVal *Node, err error) {
	if prob == 0.0 {
		return x, nil
	}

	var low, high float64
	if prob < 0 {
		low = prob
		high = -prob
	} else {
		low = -prob
		high = prob
	}

	var dt Dtype
	if dt, err = dtypeOf(x.t); err != nil {
		return nil, errors.Wrap(err, dtypeOfFail)
	}

	m := UniformRandomNode(x.g, dt, low, high, x.shape...)
	if retVal, err = HadamardProd(x, m); err != nil {
		return nil, errors.Wrap(err, mulFail)
	}

	var v Value
	switch dt {
	case Float64:
		v, _ = anyToScalar(1.0 / prob)
	case Float32:
		v, _ = anyToScalar(float32(1.0 / prob))
	default:
		// TODO: use errors package for this panic?
		panic(fmt.Sprintf("Dtype %v not yet implemented for dropout", dt))
	}

	c := NewConstant(v)

	return HadamardDiv(retVal, c)
}

// Rectify is a convenience function for creating rectified linear units activation functions.
// This function uses >=, which is the canonical version. If you want to use >, you can create
// your own by just following this.
func Rectify(x *Node) (retVal *Node, err error) {
	var zero *Node
	var dt Dtype

	// which zero to use?
	if dt, err = dtypeOf(x.t); err != nil {
		return nil, errors.Wrap(err, dtypeOfFail)
	}
	switch dt {
	case Float64:
		zero = zerof64
	case Float32:
		zero = zerof32
	default:
		return nil, errors.Errorf(nyiFail, "ReLu", dt)
	}

	cmp := newElemBinOp(gteOpType, x, zero)
	cmp.retSame = true

	if retVal, err = applyOp(cmp, x); err != nil {
		return nil, errors.Wrap(err, applyOpFail)
	}

	return HadamardProd(x, retVal)
}
