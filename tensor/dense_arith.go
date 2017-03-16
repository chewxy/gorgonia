package tensor

import (
	"math"
	"math/cmplx"
	"reflect"

	"github.com/chewxy/math32"
	"github.com/pkg/errors"
)

/*
GENERATED FILE. DO NOT EDIT
*/

func prepBinaryDense(a, b *Dense, opts ...FuncOpt) (reuse *Dense, safe, toReuse, incr bool, err error) {
	if !isNumber(a.t) && !isNumber(b.t) {
		err = noopError{}
		return
	}
	if a.t.Kind() != b.t.Kind() {
		err = errors.Errorf(typeMismatch, a.t, b.t)
		return
	}
	if !a.Shape().Eq(b.Shape()) {
		err = errors.Errorf(shapeMismatch, b.Shape(), a.Shape())
		return
	}

	fo := parseFuncOpts(opts...)
	reuseT, incr := fo.incrReuse()
	safe = fo.safe()
	toReuse = reuseT != nil

	if toReuse {
		if reuse, err = getDense(reuseT); err != nil {
			err = errors.Wrapf(err, "Cannot reuse a different type of Tensor in a *Dense-Scalar operation")
			return
		}

		if reuse.t.Kind() != a.t.Kind() {
			err = errors.Errorf(typeMismatch, a.t, reuse.t)
			err = errors.Wrapf(err, "Cannot use reuse")
			return
		}

		if reuse.len() != a.Shape().TotalSize() {
			err = errors.Errorf(shapeMismatch, reuse.Shape(), a.Shape())
			err = errors.Wrapf(err, "Cannot use reuse: shape mismatch")
			return
		}
	}
	return
}

func prepUnaryDense(a *Dense, opts ...FuncOpt) (reuse *Dense, safe, toReuse, incr bool, err error) {
	if !isNumber(a.t) {
		err = noopError{}
		return
	}

	fo := parseFuncOpts(opts...)
	reuseT, incr := fo.incrReuse()
	safe = fo.safe()
	toReuse = reuseT != nil

	if toReuse {
		if reuse, err = getDense(reuseT); err != nil {
			err = errors.Wrapf(err, "Cannot reuse a different type of Tensor in a *Dense-Scalar operation")
			return
		}

		if reuse.t.Kind() != a.t.Kind() {
			err = errors.Errorf(typeMismatch, a.t, reuse.t)
			err = errors.Wrapf(err, "Cannot use reuse")
			return
		}

		if reuse.len() != a.Shape().TotalSize() {
			err = errors.Errorf(shapeMismatch, reuse.Shape(), a.Shape())
			err = errors.Wrapf(err, "Cannot use reuse: shape mismatch")
			return
		}
	}
	return
}

/* Add */

// Add performs the operation on another *Dense. It takes a list of FuncOpts.
func (t *Dense) Add(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepBinaryDense(t, other, opts...)
	if err != nil {
		return nil, err
	}

	var it, ot *FlatIterator
	if t.IsMaterializable() {
		it = NewFlatIterator(t.AP)
	}
	if other.IsMaterializable() {
		ot = NewFlatIterator(other.AP)
	}
	switch {
	case incr:
		// when incr returned, the reuse is the *Dense to be incremented
		retVal = reuse
		switch reuse.t.Kind() {
		case reflect.Int:
			data := reuse.ints()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI(i) + other.getI(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI(i) + other.getI(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI(i) + other.getI(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI(i) + other.getI(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddI(t.ints(), other.ints(), reuse.ints())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI(i) + other.getI(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI(i) + other.getI(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI(i) + other.getI(j)
						incrI++
					}
				}
			}
		case reflect.Int8:
			data := reuse.int8s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI8(i) + other.getI8(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI8(i) + other.getI8(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI8(i) + other.getI8(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI8(i) + other.getI8(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddI8(t.int8s(), other.int8s(), reuse.int8s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI8(i) + other.getI8(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI8(i) + other.getI8(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI8(i) + other.getI8(j)
						incrI++
					}
				}
			}
		case reflect.Int16:
			data := reuse.int16s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI16(i) + other.getI16(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI16(i) + other.getI16(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI16(i) + other.getI16(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI16(i) + other.getI16(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddI16(t.int16s(), other.int16s(), reuse.int16s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI16(i) + other.getI16(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI16(i) + other.getI16(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI16(i) + other.getI16(j)
						incrI++
					}
				}
			}
		case reflect.Int32:
			data := reuse.int32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI32(i) + other.getI32(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI32(i) + other.getI32(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI32(i) + other.getI32(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI32(i) + other.getI32(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddI32(t.int32s(), other.int32s(), reuse.int32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI32(i) + other.getI32(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI32(i) + other.getI32(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI32(i) + other.getI32(j)
						incrI++
					}
				}
			}
		case reflect.Int64:
			data := reuse.int64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI64(i) + other.getI64(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI64(i) + other.getI64(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI64(i) + other.getI64(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI64(i) + other.getI64(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddI64(t.int64s(), other.int64s(), reuse.int64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI64(i) + other.getI64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI64(i) + other.getI64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI64(i) + other.getI64(j)
						incrI++
					}
				}
			}
		case reflect.Uint:
			data := reuse.uints()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU(i) + other.getU(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU(i) + other.getU(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU(i) + other.getU(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU(i) + other.getU(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddU(t.uints(), other.uints(), reuse.uints())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU(i) + other.getU(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU(i) + other.getU(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU(i) + other.getU(j)
						incrI++
					}
				}
			}
		case reflect.Uint8:
			data := reuse.uint8s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU8(i) + other.getU8(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU8(i) + other.getU8(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU8(i) + other.getU8(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU8(i) + other.getU8(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddU8(t.uint8s(), other.uint8s(), reuse.uint8s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU8(i) + other.getU8(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU8(i) + other.getU8(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU8(i) + other.getU8(j)
						incrI++
					}
				}
			}
		case reflect.Uint16:
			data := reuse.uint16s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU16(i) + other.getU16(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU16(i) + other.getU16(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU16(i) + other.getU16(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU16(i) + other.getU16(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddU16(t.uint16s(), other.uint16s(), reuse.uint16s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU16(i) + other.getU16(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU16(i) + other.getU16(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU16(i) + other.getU16(j)
						incrI++
					}
				}
			}
		case reflect.Uint32:
			data := reuse.uint32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU32(i) + other.getU32(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU32(i) + other.getU32(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU32(i) + other.getU32(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU32(i) + other.getU32(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddU32(t.uint32s(), other.uint32s(), reuse.uint32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU32(i) + other.getU32(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU32(i) + other.getU32(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU32(i) + other.getU32(j)
						incrI++
					}
				}
			}
		case reflect.Uint64:
			data := reuse.uint64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU64(i) + other.getU64(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU64(i) + other.getU64(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU64(i) + other.getU64(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU64(i) + other.getU64(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddU64(t.uint64s(), other.uint64s(), reuse.uint64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU64(i) + other.getU64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU64(i) + other.getU64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU64(i) + other.getU64(j)
						incrI++
					}
				}
			}
		case reflect.Float32:
			data := reuse.float32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getF32(i) + other.getF32(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF32(i) + other.getF32(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF32(i) + other.getF32(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getF32(i) + other.getF32(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddF32(t.float32s(), other.float32s(), reuse.float32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getF32(i) + other.getF32(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getF32(i) + other.getF32(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF32(i) + other.getF32(j)
						incrI++
					}
				}
			}
		case reflect.Float64:
			data := reuse.float64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getF64(i) + other.getF64(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF64(i) + other.getF64(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF64(i) + other.getF64(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getF64(i) + other.getF64(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddF64(t.float64s(), other.float64s(), reuse.float64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getF64(i) + other.getF64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getF64(i) + other.getF64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF64(i) + other.getF64(j)
						incrI++
					}
				}
			}
		case reflect.Complex64:
			data := reuse.complex64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getC64(i) + other.getC64(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC64(i) + other.getC64(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC64(i) + other.getC64(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getC64(i) + other.getC64(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddC64(t.complex64s(), other.complex64s(), reuse.complex64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getC64(i) + other.getC64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getC64(i) + other.getC64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC64(i) + other.getC64(j)
						incrI++
					}
				}
			}
		case reflect.Complex128:
			data := reuse.complex128s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getC128(i) + other.getC128(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC128(i) + other.getC128(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC128(i) + other.getC128(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getC128(i) + other.getC128(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecAddC128(t.complex128s(), other.complex128s(), reuse.complex128s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getC128(i) + other.getC128(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getC128(i) + other.getC128(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC128(i) + other.getC128(j)
						incrI++
					}
				}
			}
		}
	case toReuse:
		if t.IsMaterializable() {
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t) // technically copyDenseIter would have done the same but it's much slower
		}
		err = reuse.add(other, nil, ot)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		err = retVal.add(other, nil, ot)
	case !safe:
		err = t.add(other, it, ot)
		retVal = t
	}
	return
}
func (t *Dense) add(other *Dense, it, ot *FlatIterator) (err error) {
	switch t.t.Kind() {
	case reflect.Int:
		tdata := t.ints()
		odata := other.ints()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddI(tdata, odata)
		}
	case reflect.Int8:
		tdata := t.int8s()
		odata := other.int8s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddI8(tdata, odata)
		}
	case reflect.Int16:
		tdata := t.int16s()
		odata := other.int16s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddI16(tdata, odata)
		}
	case reflect.Int32:
		tdata := t.int32s()
		odata := other.int32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddI32(tdata, odata)
		}
	case reflect.Int64:
		tdata := t.int64s()
		odata := other.int64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddI64(tdata, odata)
		}
	case reflect.Uint:
		tdata := t.uints()
		odata := other.uints()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddU(tdata, odata)
		}
	case reflect.Uint8:
		tdata := t.uint8s()
		odata := other.uint8s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddU8(tdata, odata)
		}
	case reflect.Uint16:
		tdata := t.uint16s()
		odata := other.uint16s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddU16(tdata, odata)
		}
	case reflect.Uint32:
		tdata := t.uint32s()
		odata := other.uint32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddU32(tdata, odata)
		}
	case reflect.Uint64:
		tdata := t.uint64s()
		odata := other.uint64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddU64(tdata, odata)
		}
	case reflect.Float32:
		tdata := t.float32s()
		odata := other.float32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddF32(tdata, odata)
		}
	case reflect.Float64:
		tdata := t.float64s()
		odata := other.float64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddF64(tdata, odata)
		}
	case reflect.Complex64:
		tdata := t.complex64s()
		odata := other.complex64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddC64(tdata, odata)
		}
	case reflect.Complex128:
		tdata := t.complex128s()
		odata := other.complex128s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] + odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] + odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] + odata[j]
				i++
			}
		default:
			vecAddC128(tdata, odata)
		}
	default:
		// TODO: Handle Number interface
	}

	return
}

/* Sub */

// Sub performs the operation on another *Dense. It takes a list of FuncOpts.
func (t *Dense) Sub(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepBinaryDense(t, other, opts...)
	if err != nil {
		return nil, err
	}

	var it, ot *FlatIterator
	if t.IsMaterializable() {
		it = NewFlatIterator(t.AP)
	}
	if other.IsMaterializable() {
		ot = NewFlatIterator(other.AP)
	}
	switch {
	case incr:
		// when incr returned, the reuse is the *Dense to be incremented
		retVal = reuse
		switch reuse.t.Kind() {
		case reflect.Int:
			data := reuse.ints()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI(i) - other.getI(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI(i) - other.getI(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI(i) - other.getI(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI(i) - other.getI(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubI(t.ints(), other.ints(), reuse.ints())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI(i) - other.getI(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI(i) - other.getI(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI(i) - other.getI(j)
						incrI++
					}
				}
			}
		case reflect.Int8:
			data := reuse.int8s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI8(i) - other.getI8(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI8(i) - other.getI8(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI8(i) - other.getI8(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI8(i) - other.getI8(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubI8(t.int8s(), other.int8s(), reuse.int8s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI8(i) - other.getI8(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI8(i) - other.getI8(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI8(i) - other.getI8(j)
						incrI++
					}
				}
			}
		case reflect.Int16:
			data := reuse.int16s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI16(i) - other.getI16(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI16(i) - other.getI16(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI16(i) - other.getI16(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI16(i) - other.getI16(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubI16(t.int16s(), other.int16s(), reuse.int16s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI16(i) - other.getI16(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI16(i) - other.getI16(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI16(i) - other.getI16(j)
						incrI++
					}
				}
			}
		case reflect.Int32:
			data := reuse.int32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI32(i) - other.getI32(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI32(i) - other.getI32(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI32(i) - other.getI32(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI32(i) - other.getI32(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubI32(t.int32s(), other.int32s(), reuse.int32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI32(i) - other.getI32(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI32(i) - other.getI32(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI32(i) - other.getI32(j)
						incrI++
					}
				}
			}
		case reflect.Int64:
			data := reuse.int64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI64(i) - other.getI64(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI64(i) - other.getI64(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI64(i) - other.getI64(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI64(i) - other.getI64(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubI64(t.int64s(), other.int64s(), reuse.int64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI64(i) - other.getI64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI64(i) - other.getI64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI64(i) - other.getI64(j)
						incrI++
					}
				}
			}
		case reflect.Uint:
			data := reuse.uints()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU(i) - other.getU(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU(i) - other.getU(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU(i) - other.getU(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU(i) - other.getU(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubU(t.uints(), other.uints(), reuse.uints())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU(i) - other.getU(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU(i) - other.getU(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU(i) - other.getU(j)
						incrI++
					}
				}
			}
		case reflect.Uint8:
			data := reuse.uint8s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU8(i) - other.getU8(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU8(i) - other.getU8(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU8(i) - other.getU8(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU8(i) - other.getU8(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubU8(t.uint8s(), other.uint8s(), reuse.uint8s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU8(i) - other.getU8(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU8(i) - other.getU8(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU8(i) - other.getU8(j)
						incrI++
					}
				}
			}
		case reflect.Uint16:
			data := reuse.uint16s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU16(i) - other.getU16(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU16(i) - other.getU16(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU16(i) - other.getU16(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU16(i) - other.getU16(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubU16(t.uint16s(), other.uint16s(), reuse.uint16s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU16(i) - other.getU16(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU16(i) - other.getU16(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU16(i) - other.getU16(j)
						incrI++
					}
				}
			}
		case reflect.Uint32:
			data := reuse.uint32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU32(i) - other.getU32(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU32(i) - other.getU32(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU32(i) - other.getU32(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU32(i) - other.getU32(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubU32(t.uint32s(), other.uint32s(), reuse.uint32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU32(i) - other.getU32(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU32(i) - other.getU32(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU32(i) - other.getU32(j)
						incrI++
					}
				}
			}
		case reflect.Uint64:
			data := reuse.uint64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU64(i) - other.getU64(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU64(i) - other.getU64(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU64(i) - other.getU64(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU64(i) - other.getU64(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubU64(t.uint64s(), other.uint64s(), reuse.uint64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU64(i) - other.getU64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU64(i) - other.getU64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU64(i) - other.getU64(j)
						incrI++
					}
				}
			}
		case reflect.Float32:
			data := reuse.float32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getF32(i) - other.getF32(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF32(i) - other.getF32(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF32(i) - other.getF32(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getF32(i) - other.getF32(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubF32(t.float32s(), other.float32s(), reuse.float32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getF32(i) - other.getF32(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getF32(i) - other.getF32(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF32(i) - other.getF32(j)
						incrI++
					}
				}
			}
		case reflect.Float64:
			data := reuse.float64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getF64(i) - other.getF64(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF64(i) - other.getF64(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF64(i) - other.getF64(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getF64(i) - other.getF64(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubF64(t.float64s(), other.float64s(), reuse.float64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getF64(i) - other.getF64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getF64(i) - other.getF64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF64(i) - other.getF64(j)
						incrI++
					}
				}
			}
		case reflect.Complex64:
			data := reuse.complex64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getC64(i) - other.getC64(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC64(i) - other.getC64(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC64(i) - other.getC64(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getC64(i) - other.getC64(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubC64(t.complex64s(), other.complex64s(), reuse.complex64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getC64(i) - other.getC64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getC64(i) - other.getC64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC64(i) - other.getC64(j)
						incrI++
					}
				}
			}
		case reflect.Complex128:
			data := reuse.complex128s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getC128(i) - other.getC128(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC128(i) - other.getC128(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC128(i) - other.getC128(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getC128(i) - other.getC128(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecSubC128(t.complex128s(), other.complex128s(), reuse.complex128s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getC128(i) - other.getC128(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getC128(i) - other.getC128(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC128(i) - other.getC128(j)
						incrI++
					}
				}
			}
		}
	case toReuse:
		if t.IsMaterializable() {
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t) // technically copyDenseIter would have done the same but it's much slower
		}
		err = reuse.sub(other, nil, ot)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		err = retVal.sub(other, nil, ot)
	case !safe:
		err = t.sub(other, it, ot)
		retVal = t
	}
	return
}
func (t *Dense) sub(other *Dense, it, ot *FlatIterator) (err error) {
	switch t.t.Kind() {
	case reflect.Int:
		tdata := t.ints()
		odata := other.ints()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubI(tdata, odata)
		}
	case reflect.Int8:
		tdata := t.int8s()
		odata := other.int8s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubI8(tdata, odata)
		}
	case reflect.Int16:
		tdata := t.int16s()
		odata := other.int16s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubI16(tdata, odata)
		}
	case reflect.Int32:
		tdata := t.int32s()
		odata := other.int32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubI32(tdata, odata)
		}
	case reflect.Int64:
		tdata := t.int64s()
		odata := other.int64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubI64(tdata, odata)
		}
	case reflect.Uint:
		tdata := t.uints()
		odata := other.uints()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubU(tdata, odata)
		}
	case reflect.Uint8:
		tdata := t.uint8s()
		odata := other.uint8s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubU8(tdata, odata)
		}
	case reflect.Uint16:
		tdata := t.uint16s()
		odata := other.uint16s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubU16(tdata, odata)
		}
	case reflect.Uint32:
		tdata := t.uint32s()
		odata := other.uint32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubU32(tdata, odata)
		}
	case reflect.Uint64:
		tdata := t.uint64s()
		odata := other.uint64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubU64(tdata, odata)
		}
	case reflect.Float32:
		tdata := t.float32s()
		odata := other.float32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubF32(tdata, odata)
		}
	case reflect.Float64:
		tdata := t.float64s()
		odata := other.float64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubF64(tdata, odata)
		}
	case reflect.Complex64:
		tdata := t.complex64s()
		odata := other.complex64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubC64(tdata, odata)
		}
	case reflect.Complex128:
		tdata := t.complex128s()
		odata := other.complex128s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] - odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] - odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] - odata[j]
				i++
			}
		default:
			vecSubC128(tdata, odata)
		}
	default:
		// TODO: Handle Number interface
	}

	return
}

/* Mul */

// Mul performs the operation on another *Dense. It takes a list of FuncOpts.
func (t *Dense) Mul(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepBinaryDense(t, other, opts...)
	if err != nil {
		return nil, err
	}

	var it, ot *FlatIterator
	if t.IsMaterializable() {
		it = NewFlatIterator(t.AP)
	}
	if other.IsMaterializable() {
		ot = NewFlatIterator(other.AP)
	}
	switch {
	case incr:
		// when incr returned, the reuse is the *Dense to be incremented
		retVal = reuse
		switch reuse.t.Kind() {
		case reflect.Int:
			data := reuse.ints()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI(i) * other.getI(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI(i) * other.getI(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI(i) * other.getI(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI(i) * other.getI(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulI(t.ints(), other.ints(), reuse.ints())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI(i) * other.getI(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI(i) * other.getI(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI(i) * other.getI(j)
						incrI++
					}
				}
			}
		case reflect.Int8:
			data := reuse.int8s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI8(i) * other.getI8(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI8(i) * other.getI8(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI8(i) * other.getI8(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI8(i) * other.getI8(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulI8(t.int8s(), other.int8s(), reuse.int8s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI8(i) * other.getI8(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI8(i) * other.getI8(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI8(i) * other.getI8(j)
						incrI++
					}
				}
			}
		case reflect.Int16:
			data := reuse.int16s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI16(i) * other.getI16(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI16(i) * other.getI16(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI16(i) * other.getI16(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI16(i) * other.getI16(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulI16(t.int16s(), other.int16s(), reuse.int16s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI16(i) * other.getI16(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI16(i) * other.getI16(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI16(i) * other.getI16(j)
						incrI++
					}
				}
			}
		case reflect.Int32:
			data := reuse.int32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI32(i) * other.getI32(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI32(i) * other.getI32(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI32(i) * other.getI32(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI32(i) * other.getI32(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulI32(t.int32s(), other.int32s(), reuse.int32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI32(i) * other.getI32(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI32(i) * other.getI32(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI32(i) * other.getI32(j)
						incrI++
					}
				}
			}
		case reflect.Int64:
			data := reuse.int64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getI64(i) * other.getI64(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI64(i) * other.getI64(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI64(i) * other.getI64(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getI64(i) * other.getI64(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulI64(t.int64s(), other.int64s(), reuse.int64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getI64(i) * other.getI64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getI64(i) * other.getI64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getI64(i) * other.getI64(j)
						incrI++
					}
				}
			}
		case reflect.Uint:
			data := reuse.uints()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU(i) * other.getU(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU(i) * other.getU(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU(i) * other.getU(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU(i) * other.getU(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulU(t.uints(), other.uints(), reuse.uints())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU(i) * other.getU(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU(i) * other.getU(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU(i) * other.getU(j)
						incrI++
					}
				}
			}
		case reflect.Uint8:
			data := reuse.uint8s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU8(i) * other.getU8(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU8(i) * other.getU8(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU8(i) * other.getU8(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU8(i) * other.getU8(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulU8(t.uint8s(), other.uint8s(), reuse.uint8s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU8(i) * other.getU8(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU8(i) * other.getU8(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU8(i) * other.getU8(j)
						incrI++
					}
				}
			}
		case reflect.Uint16:
			data := reuse.uint16s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU16(i) * other.getU16(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU16(i) * other.getU16(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU16(i) * other.getU16(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU16(i) * other.getU16(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulU16(t.uint16s(), other.uint16s(), reuse.uint16s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU16(i) * other.getU16(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU16(i) * other.getU16(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU16(i) * other.getU16(j)
						incrI++
					}
				}
			}
		case reflect.Uint32:
			data := reuse.uint32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU32(i) * other.getU32(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU32(i) * other.getU32(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU32(i) * other.getU32(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU32(i) * other.getU32(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulU32(t.uint32s(), other.uint32s(), reuse.uint32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU32(i) * other.getU32(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU32(i) * other.getU32(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU32(i) * other.getU32(j)
						incrI++
					}
				}
			}
		case reflect.Uint64:
			data := reuse.uint64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getU64(i) * other.getU64(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU64(i) * other.getU64(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU64(i) * other.getU64(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getU64(i) * other.getU64(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulU64(t.uint64s(), other.uint64s(), reuse.uint64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getU64(i) * other.getU64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getU64(i) * other.getU64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getU64(i) * other.getU64(j)
						incrI++
					}
				}
			}
		case reflect.Float32:
			data := reuse.float32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getF32(i) * other.getF32(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF32(i) * other.getF32(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF32(i) * other.getF32(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getF32(i) * other.getF32(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulF32(t.float32s(), other.float32s(), reuse.float32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getF32(i) * other.getF32(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getF32(i) * other.getF32(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF32(i) * other.getF32(j)
						incrI++
					}
				}
			}
		case reflect.Float64:
			data := reuse.float64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getF64(i) * other.getF64(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF64(i) * other.getF64(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF64(i) * other.getF64(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getF64(i) * other.getF64(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulF64(t.float64s(), other.float64s(), reuse.float64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getF64(i) * other.getF64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getF64(i) * other.getF64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF64(i) * other.getF64(j)
						incrI++
					}
				}
			}
		case reflect.Complex64:
			data := reuse.complex64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getC64(i) * other.getC64(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC64(i) * other.getC64(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC64(i) * other.getC64(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getC64(i) * other.getC64(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulC64(t.complex64s(), other.complex64s(), reuse.complex64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getC64(i) * other.getC64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getC64(i) * other.getC64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC64(i) * other.getC64(j)
						incrI++
					}
				}
			}
		case reflect.Complex128:
			data := reuse.complex128s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getC128(i) * other.getC128(j)
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC128(i) * other.getC128(j)
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC128(i) * other.getC128(j)
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getC128(i) * other.getC128(j)
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecMulC128(t.complex128s(), other.complex128s(), reuse.complex128s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getC128(i) * other.getC128(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getC128(i) * other.getC128(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC128(i) * other.getC128(j)
						incrI++
					}
				}
			}
		}
	case toReuse:
		if t.IsMaterializable() {
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t) // technically copyDenseIter would have done the same but it's much slower
		}
		err = reuse.mul(other, nil, ot)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		err = retVal.mul(other, nil, ot)
	case !safe:
		err = t.mul(other, it, ot)
		retVal = t
	}
	return
}
func (t *Dense) mul(other *Dense, it, ot *FlatIterator) (err error) {
	switch t.t.Kind() {
	case reflect.Int:
		tdata := t.ints()
		odata := other.ints()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulI(tdata, odata)
		}
	case reflect.Int8:
		tdata := t.int8s()
		odata := other.int8s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulI8(tdata, odata)
		}
	case reflect.Int16:
		tdata := t.int16s()
		odata := other.int16s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulI16(tdata, odata)
		}
	case reflect.Int32:
		tdata := t.int32s()
		odata := other.int32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulI32(tdata, odata)
		}
	case reflect.Int64:
		tdata := t.int64s()
		odata := other.int64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulI64(tdata, odata)
		}
	case reflect.Uint:
		tdata := t.uints()
		odata := other.uints()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulU(tdata, odata)
		}
	case reflect.Uint8:
		tdata := t.uint8s()
		odata := other.uint8s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulU8(tdata, odata)
		}
	case reflect.Uint16:
		tdata := t.uint16s()
		odata := other.uint16s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulU16(tdata, odata)
		}
	case reflect.Uint32:
		tdata := t.uint32s()
		odata := other.uint32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulU32(tdata, odata)
		}
	case reflect.Uint64:
		tdata := t.uint64s()
		odata := other.uint64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulU64(tdata, odata)
		}
	case reflect.Float32:
		tdata := t.float32s()
		odata := other.float32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulF32(tdata, odata)
		}
	case reflect.Float64:
		tdata := t.float64s()
		odata := other.float64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulF64(tdata, odata)
		}
	case reflect.Complex64:
		tdata := t.complex64s()
		odata := other.complex64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulC64(tdata, odata)
		}
	case reflect.Complex128:
		tdata := t.complex128s()
		odata := other.complex128s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] * odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] * odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] * odata[j]
				i++
			}
		default:
			vecMulC128(tdata, odata)
		}
	default:
		// TODO: Handle Number interface
	}

	return
}

/* Div */

// Div performs the operation on another *Dense. It takes a list of FuncOpts.
func (t *Dense) Div(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepBinaryDense(t, other, opts...)
	if err != nil {
		return nil, err
	}

	var errs errorIndices
	var it, ot *FlatIterator
	if t.IsMaterializable() {
		it = NewFlatIterator(t.AP)
	}
	if other.IsMaterializable() {
		ot = NewFlatIterator(other.AP)
	}
	switch {
	case incr:
		// when incr returned, the reuse is the *Dense to be incremented
		retVal = reuse
		switch reuse.t.Kind() {
		case reflect.Int:
			data := reuse.ints()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						if other.getI(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI(i) / other.getI(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI(i) / other.getI(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI(i) / other.getI(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						if other.getI(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI(i) / other.getI(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivI(t.ints(), other.ints(), reuse.ints())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						if other.getI(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI(i) / other.getI(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						if other.getI(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI(i) / other.getI(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI(i) / other.getI(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Int8:
			data := reuse.int8s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						if other.getI8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI8(i) / other.getI8(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI8(i) / other.getI8(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI8(i) / other.getI8(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						if other.getI8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI8(i) / other.getI8(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivI8(t.int8s(), other.int8s(), reuse.int8s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						if other.getI8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI8(i) / other.getI8(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						if other.getI8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI8(i) / other.getI8(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI8(i) / other.getI8(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Int16:
			data := reuse.int16s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						if other.getI16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI16(i) / other.getI16(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI16(i) / other.getI16(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI16(i) / other.getI16(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						if other.getI16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI16(i) / other.getI16(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivI16(t.int16s(), other.int16s(), reuse.int16s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						if other.getI16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI16(i) / other.getI16(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						if other.getI16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI16(i) / other.getI16(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI16(i) / other.getI16(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Int32:
			data := reuse.int32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						if other.getI32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI32(i) / other.getI32(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI32(i) / other.getI32(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI32(i) / other.getI32(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						if other.getI32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI32(i) / other.getI32(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivI32(t.int32s(), other.int32s(), reuse.int32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						if other.getI32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI32(i) / other.getI32(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						if other.getI32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI32(i) / other.getI32(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI32(i) / other.getI32(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Int64:
			data := reuse.int64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						if other.getI64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI64(i) / other.getI64(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI64(i) / other.getI64(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI64(i) / other.getI64(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						if other.getI64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI64(i) / other.getI64(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivI64(t.int64s(), other.int64s(), reuse.int64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						if other.getI64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI64(i) / other.getI64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						if other.getI64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI64(i) / other.getI64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getI64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getI64(i) / other.getI64(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Uint:
			data := reuse.uints()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						if other.getU(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU(i) / other.getU(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU(i) / other.getU(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU(i) / other.getU(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						if other.getU(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU(i) / other.getU(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivU(t.uints(), other.uints(), reuse.uints())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						if other.getU(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU(i) / other.getU(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						if other.getU(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU(i) / other.getU(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU(i) / other.getU(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Uint8:
			data := reuse.uint8s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						if other.getU8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU8(i) / other.getU8(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU8(i) / other.getU8(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU8(i) / other.getU8(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						if other.getU8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU8(i) / other.getU8(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivU8(t.uint8s(), other.uint8s(), reuse.uint8s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						if other.getU8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU8(i) / other.getU8(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						if other.getU8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU8(i) / other.getU8(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU8(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU8(i) / other.getU8(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Uint16:
			data := reuse.uint16s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						if other.getU16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU16(i) / other.getU16(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU16(i) / other.getU16(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU16(i) / other.getU16(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						if other.getU16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU16(i) / other.getU16(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivU16(t.uint16s(), other.uint16s(), reuse.uint16s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						if other.getU16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU16(i) / other.getU16(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						if other.getU16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU16(i) / other.getU16(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU16(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU16(i) / other.getU16(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Uint32:
			data := reuse.uint32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						if other.getU32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU32(i) / other.getU32(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU32(i) / other.getU32(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU32(i) / other.getU32(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						if other.getU32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU32(i) / other.getU32(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivU32(t.uint32s(), other.uint32s(), reuse.uint32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						if other.getU32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU32(i) / other.getU32(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						if other.getU32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU32(i) / other.getU32(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU32(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU32(i) / other.getU32(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Uint64:
			data := reuse.uint64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						if other.getU64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU64(i) / other.getU64(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU64(i) / other.getU64(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU64(i) / other.getU64(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						if other.getU64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU64(i) / other.getU64(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivU64(t.uint64s(), other.uint64s(), reuse.uint64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						if other.getU64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU64(i) / other.getU64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						if other.getU64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU64(i) / other.getU64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if other.getU64(j) == 0 {
							errs = append(errs, j)
							continue
						}
						data[incrI] += t.getU64(i) / other.getU64(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Float32:
			data := reuse.float32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getF32(i) / other.getF32(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF32(i) / other.getF32(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF32(i) / other.getF32(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getF32(i) / other.getF32(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivF32(t.float32s(), other.float32s(), reuse.float32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getF32(i) / other.getF32(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getF32(i) / other.getF32(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF32(i) / other.getF32(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Float64:
			data := reuse.float64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getF64(i) / other.getF64(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF64(i) / other.getF64(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF64(i) / other.getF64(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getF64(i) / other.getF64(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivF64(t.float64s(), other.float64s(), reuse.float64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getF64(i) / other.getF64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getF64(i) / other.getF64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getF64(i) / other.getF64(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Complex64:
			data := reuse.complex64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getC64(i) / other.getC64(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC64(i) / other.getC64(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC64(i) / other.getC64(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getC64(i) / other.getC64(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivC64(t.complex64s(), other.complex64s(), reuse.complex64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getC64(i) / other.getC64(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getC64(i) / other.getC64(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC64(i) / other.getC64(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		case reflect.Complex128:
			data := reuse.complex128s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += t.getC128(i) / other.getC128(j)
						i++
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC128(i) / other.getC128(j)
						j++
					}
					if err != nil {
						return
					}
					err = errs
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC128(i) / other.getC128(j)
						i++
					}
					if err != nil {
						return
					}
					err = errs
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += t.getC128(i) / other.getC128(j)
					}
					if err != nil {
						return
					}
					err = errs
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecDivC128(t.complex128s(), other.complex128s(), reuse.complex128s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += t.getC128(i) / other.getC128(j)
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += t.getC128(i) / other.getC128(j)
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += t.getC128(i) / other.getC128(j)
						incrI++
					}
					if err != nil {
						return
					}
					err = errs
				}
			}
		}
		if errs != nil {
			err = errs
		}
	case toReuse:
		if t.IsMaterializable() {
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t) // technically copyDenseIter would have done the same but it's much slower
		}
		err = reuse.div(other, nil, ot)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		err = retVal.div(other, nil, ot)
	case !safe:
		err = t.div(other, it, ot)
		retVal = t
	}
	return
}
func (t *Dense) div(other *Dense, it, ot *FlatIterator) (err error) {
	var errs errorIndices
	switch t.t.Kind() {
	case reflect.Int:
		tdata := t.ints()
		odata := other.ints()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivI(tdata, odata)
		}
	case reflect.Int8:
		tdata := t.int8s()
		odata := other.int8s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivI8(tdata, odata)
		}
	case reflect.Int16:
		tdata := t.int16s()
		odata := other.int16s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivI16(tdata, odata)
		}
	case reflect.Int32:
		tdata := t.int32s()
		odata := other.int32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivI32(tdata, odata)
		}
	case reflect.Int64:
		tdata := t.int64s()
		odata := other.int64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivI64(tdata, odata)
		}
	case reflect.Uint:
		tdata := t.uints()
		odata := other.uints()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivU(tdata, odata)
		}
	case reflect.Uint8:
		tdata := t.uint8s()
		odata := other.uint8s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivU8(tdata, odata)
		}
	case reflect.Uint16:
		tdata := t.uint16s()
		odata := other.uint16s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivU16(tdata, odata)
		}
	case reflect.Uint32:
		tdata := t.uint32s()
		odata := other.uint32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivU32(tdata, odata)
		}
	case reflect.Uint64:
		tdata := t.uint64s()
		odata := other.uint64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				if odata[j] == 0 {
					errs = append(errs, j)
					continue
				}
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivU64(tdata, odata)
		}
	case reflect.Float32:
		tdata := t.float32s()
		odata := other.float32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivF32(tdata, odata)
		}
	case reflect.Float64:
		tdata := t.float64s()
		odata := other.float64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivF64(tdata, odata)
		}
	case reflect.Complex64:
		tdata := t.complex64s()
		odata := other.complex64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivC64(tdata, odata)
		}
	case reflect.Complex128:
		tdata := t.complex128s()
		odata := other.complex128s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = tdata[i] / odata[j]
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = tdata[i] / odata[j]
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = tdata[i] / odata[j]
				i++
			}
		default:
			vecDivC128(tdata, odata)
		}
	default:
		// TODO: Handle Number interface
	}

	if err != nil {
		return
	}

	if errs != nil {
		err = errs
	}
	return
}

/* Pow */

// Pow performs the operation on another *Dense. It takes a list of FuncOpts.
func (t *Dense) Pow(other *Dense, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepBinaryDense(t, other, opts...)
	if err != nil {
		return nil, err
	}

	var it, ot *FlatIterator
	if t.IsMaterializable() {
		it = NewFlatIterator(t.AP)
	}
	if other.IsMaterializable() {
		ot = NewFlatIterator(other.AP)
	}
	switch {
	case incr:
		// when incr returned, the reuse is the *Dense to be incremented
		retVal = reuse
		switch reuse.t.Kind() {
		case reflect.Int:
			data := reuse.ints()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += int(math.Pow(float64(t.getI(i)), float64(other.getI(j))))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int(math.Pow(float64(t.getI(i)), float64(other.getI(j))))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int(math.Pow(float64(t.getI(i)), float64(other.getI(j))))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += int(math.Pow(float64(t.getI(i)), float64(other.getI(j))))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowI(t.ints(), other.ints(), reuse.ints())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += int(math.Pow(float64(t.getI(i)), float64(other.getI(j))))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += int(math.Pow(float64(t.getI(i)), float64(other.getI(j))))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int(math.Pow(float64(t.getI(i)), float64(other.getI(j))))
						incrI++
					}
				}
			}
		case reflect.Int8:
			data := reuse.int8s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += int8(math.Pow(float64(t.getI8(i)), float64(other.getI8(j))))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int8(math.Pow(float64(t.getI8(i)), float64(other.getI8(j))))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int8(math.Pow(float64(t.getI8(i)), float64(other.getI8(j))))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += int8(math.Pow(float64(t.getI8(i)), float64(other.getI8(j))))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowI8(t.int8s(), other.int8s(), reuse.int8s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += int8(math.Pow(float64(t.getI8(i)), float64(other.getI8(j))))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += int8(math.Pow(float64(t.getI8(i)), float64(other.getI8(j))))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int8(math.Pow(float64(t.getI8(i)), float64(other.getI8(j))))
						incrI++
					}
				}
			}
		case reflect.Int16:
			data := reuse.int16s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += int16(math.Pow(float64(t.getI16(i)), float64(other.getI16(j))))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int16(math.Pow(float64(t.getI16(i)), float64(other.getI16(j))))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int16(math.Pow(float64(t.getI16(i)), float64(other.getI16(j))))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += int16(math.Pow(float64(t.getI16(i)), float64(other.getI16(j))))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowI16(t.int16s(), other.int16s(), reuse.int16s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += int16(math.Pow(float64(t.getI16(i)), float64(other.getI16(j))))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += int16(math.Pow(float64(t.getI16(i)), float64(other.getI16(j))))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int16(math.Pow(float64(t.getI16(i)), float64(other.getI16(j))))
						incrI++
					}
				}
			}
		case reflect.Int32:
			data := reuse.int32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += int32(math.Pow(float64(t.getI32(i)), float64(other.getI32(j))))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int32(math.Pow(float64(t.getI32(i)), float64(other.getI32(j))))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int32(math.Pow(float64(t.getI32(i)), float64(other.getI32(j))))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += int32(math.Pow(float64(t.getI32(i)), float64(other.getI32(j))))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowI32(t.int32s(), other.int32s(), reuse.int32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += int32(math.Pow(float64(t.getI32(i)), float64(other.getI32(j))))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += int32(math.Pow(float64(t.getI32(i)), float64(other.getI32(j))))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int32(math.Pow(float64(t.getI32(i)), float64(other.getI32(j))))
						incrI++
					}
				}
			}
		case reflect.Int64:
			data := reuse.int64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += int64(math.Pow(float64(t.getI64(i)), float64(other.getI64(j))))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int64(math.Pow(float64(t.getI64(i)), float64(other.getI64(j))))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int64(math.Pow(float64(t.getI64(i)), float64(other.getI64(j))))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += int64(math.Pow(float64(t.getI64(i)), float64(other.getI64(j))))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowI64(t.int64s(), other.int64s(), reuse.int64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += int64(math.Pow(float64(t.getI64(i)), float64(other.getI64(j))))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += int64(math.Pow(float64(t.getI64(i)), float64(other.getI64(j))))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += int64(math.Pow(float64(t.getI64(i)), float64(other.getI64(j))))
						incrI++
					}
				}
			}
		case reflect.Uint:
			data := reuse.uints()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += uint(math.Pow(float64(t.getU(i)), float64(other.getU(j))))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint(math.Pow(float64(t.getU(i)), float64(other.getU(j))))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint(math.Pow(float64(t.getU(i)), float64(other.getU(j))))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += uint(math.Pow(float64(t.getU(i)), float64(other.getU(j))))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowU(t.uints(), other.uints(), reuse.uints())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += uint(math.Pow(float64(t.getU(i)), float64(other.getU(j))))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += uint(math.Pow(float64(t.getU(i)), float64(other.getU(j))))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint(math.Pow(float64(t.getU(i)), float64(other.getU(j))))
						incrI++
					}
				}
			}
		case reflect.Uint8:
			data := reuse.uint8s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += uint8(math.Pow(float64(t.getU8(i)), float64(other.getU8(j))))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint8(math.Pow(float64(t.getU8(i)), float64(other.getU8(j))))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint8(math.Pow(float64(t.getU8(i)), float64(other.getU8(j))))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += uint8(math.Pow(float64(t.getU8(i)), float64(other.getU8(j))))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowU8(t.uint8s(), other.uint8s(), reuse.uint8s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += uint8(math.Pow(float64(t.getU8(i)), float64(other.getU8(j))))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += uint8(math.Pow(float64(t.getU8(i)), float64(other.getU8(j))))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint8(math.Pow(float64(t.getU8(i)), float64(other.getU8(j))))
						incrI++
					}
				}
			}
		case reflect.Uint16:
			data := reuse.uint16s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += uint16(math.Pow(float64(t.getU16(i)), float64(other.getU16(j))))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint16(math.Pow(float64(t.getU16(i)), float64(other.getU16(j))))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint16(math.Pow(float64(t.getU16(i)), float64(other.getU16(j))))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += uint16(math.Pow(float64(t.getU16(i)), float64(other.getU16(j))))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowU16(t.uint16s(), other.uint16s(), reuse.uint16s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += uint16(math.Pow(float64(t.getU16(i)), float64(other.getU16(j))))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += uint16(math.Pow(float64(t.getU16(i)), float64(other.getU16(j))))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint16(math.Pow(float64(t.getU16(i)), float64(other.getU16(j))))
						incrI++
					}
				}
			}
		case reflect.Uint32:
			data := reuse.uint32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += uint32(math.Pow(float64(t.getU32(i)), float64(other.getU32(j))))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint32(math.Pow(float64(t.getU32(i)), float64(other.getU32(j))))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint32(math.Pow(float64(t.getU32(i)), float64(other.getU32(j))))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += uint32(math.Pow(float64(t.getU32(i)), float64(other.getU32(j))))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowU32(t.uint32s(), other.uint32s(), reuse.uint32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += uint32(math.Pow(float64(t.getU32(i)), float64(other.getU32(j))))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += uint32(math.Pow(float64(t.getU32(i)), float64(other.getU32(j))))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint32(math.Pow(float64(t.getU32(i)), float64(other.getU32(j))))
						incrI++
					}
				}
			}
		case reflect.Uint64:
			data := reuse.uint64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += uint64(math.Pow(float64(t.getU64(i)), float64(other.getU64(j))))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint64(math.Pow(float64(t.getU64(i)), float64(other.getU64(j))))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint64(math.Pow(float64(t.getU64(i)), float64(other.getU64(j))))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += uint64(math.Pow(float64(t.getU64(i)), float64(other.getU64(j))))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowU64(t.uint64s(), other.uint64s(), reuse.uint64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += uint64(math.Pow(float64(t.getU64(i)), float64(other.getU64(j))))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += uint64(math.Pow(float64(t.getU64(i)), float64(other.getU64(j))))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += uint64(math.Pow(float64(t.getU64(i)), float64(other.getU64(j))))
						incrI++
					}
				}
			}
		case reflect.Float32:
			data := reuse.float32s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += math32.Pow(t.getF32(i), other.getF32(j))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += math32.Pow(t.getF32(i), other.getF32(j))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += math32.Pow(t.getF32(i), other.getF32(j))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += math32.Pow(t.getF32(i), other.getF32(j))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowF32(t.float32s(), other.float32s(), reuse.float32s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += math32.Pow(t.getF32(i), other.getF32(j))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += math32.Pow(t.getF32(i), other.getF32(j))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += math32.Pow(t.getF32(i), other.getF32(j))
						incrI++
					}
				}
			}
		case reflect.Float64:
			data := reuse.float64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += math.Pow(t.getF64(i), other.getF64(j))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += math.Pow(t.getF64(i), other.getF64(j))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += math.Pow(t.getF64(i), other.getF64(j))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += math.Pow(t.getF64(i), other.getF64(j))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowF64(t.float64s(), other.float64s(), reuse.float64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += math.Pow(t.getF64(i), other.getF64(j))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += math.Pow(t.getF64(i), other.getF64(j))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += math.Pow(t.getF64(i), other.getF64(j))
						incrI++
					}
				}
			}
		case reflect.Complex64:
			data := reuse.complex64s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += complex64(cmplx.Pow(complex128(t.getC64(i)), complex128(other.getC64(j))))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += complex64(cmplx.Pow(complex128(t.getC64(i)), complex128(other.getC64(j))))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += complex64(cmplx.Pow(complex128(t.getC64(i)), complex128(other.getC64(j))))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += complex64(cmplx.Pow(complex128(t.getC64(i)), complex128(other.getC64(j))))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowC64(t.complex64s(), other.complex64s(), reuse.complex64s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += complex64(cmplx.Pow(complex128(t.getC64(i)), complex128(other.getC64(j))))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += complex64(cmplx.Pow(complex128(t.getC64(i)), complex128(other.getC64(j))))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += complex64(cmplx.Pow(complex128(t.getC64(i)), complex128(other.getC64(j))))
						incrI++
					}
				}
			}
		case reflect.Complex128:
			data := reuse.complex128s()
			switch {
			case reuse.IsMaterializable():
				incrIter := NewFlatIterator(reuse.AP)
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					for incrI, err = incrIter.Next(); err == nil; incrI, err = incrIter.Next() {
						data[incrI] += cmplx.Pow(t.getC128(i), other.getC128(j))
						i++
						j++
					}
				case it != nil && ot == nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += cmplx.Pow(t.getC128(i), other.getC128(j))
						j++
					}
				case it == nil && ot != nil:
					for {
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += cmplx.Pow(t.getC128(i), other.getC128(j))
						i++
					}
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if incrI, err = incrIter.Next(); err != nil {
							err = handleNoOp(err)
							break
						}

						data[incrI] += cmplx.Pow(t.getC128(i), other.getC128(j))
					}
				}
			case !reuse.IsMaterializable():
				var i, j, incrI int
				switch {
				case it == nil && ot == nil:
					err = incrVecPowC128(t.complex128s(), other.complex128s(), reuse.complex128s())
				case it != nil && ot == nil:
					for i, err = it.Next(); err == nil; i, err = it.Next() {
						data[incrI] += cmplx.Pow(t.getC128(i), other.getC128(j))
						j++
						incrI++
					}
					err = handleNoOp(err)
				case it == nil && ot != nil:
					for j, err = ot.Next(); err == nil; j, err = ot.Next() {
						data[incrI] += cmplx.Pow(t.getC128(i), other.getC128(j))
						i++
						incrI++
					}
					err = handleNoOp(err)
				case it != nil && ot != nil:
					for {
						if i, err = it.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						if j, err = ot.Next(); err != nil {
							err = handleNoOp(err)
							break
						}
						data[incrI] += cmplx.Pow(t.getC128(i), other.getC128(j))
						incrI++
					}
				}
			}
		}
	case toReuse:
		if t.IsMaterializable() {
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t) // technically copyDenseIter would have done the same but it's much slower
		}
		err = reuse.pow(other, nil, ot)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		err = retVal.pow(other, nil, ot)
	case !safe:
		err = t.pow(other, it, ot)
		retVal = t
	}
	return
}
func (t *Dense) pow(other *Dense, it, ot *FlatIterator) (err error) {
	switch t.t.Kind() {
	case reflect.Int:
		tdata := t.ints()
		odata := other.ints()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = int(math.Pow(float64(tdata[i]), float64(odata[j])))
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = int(math.Pow(float64(tdata[i]), float64(odata[j])))
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = int(math.Pow(float64(tdata[i]), float64(odata[j])))
				i++
			}
		default:
			vecPowI(tdata, odata)
		}
	case reflect.Int8:
		tdata := t.int8s()
		odata := other.int8s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = int8(math.Pow(float64(tdata[i]), float64(odata[j])))
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = int8(math.Pow(float64(tdata[i]), float64(odata[j])))
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = int8(math.Pow(float64(tdata[i]), float64(odata[j])))
				i++
			}
		default:
			vecPowI8(tdata, odata)
		}
	case reflect.Int16:
		tdata := t.int16s()
		odata := other.int16s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = int16(math.Pow(float64(tdata[i]), float64(odata[j])))
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = int16(math.Pow(float64(tdata[i]), float64(odata[j])))
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = int16(math.Pow(float64(tdata[i]), float64(odata[j])))
				i++
			}
		default:
			vecPowI16(tdata, odata)
		}
	case reflect.Int32:
		tdata := t.int32s()
		odata := other.int32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = int32(math.Pow(float64(tdata[i]), float64(odata[j])))
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = int32(math.Pow(float64(tdata[i]), float64(odata[j])))
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = int32(math.Pow(float64(tdata[i]), float64(odata[j])))
				i++
			}
		default:
			vecPowI32(tdata, odata)
		}
	case reflect.Int64:
		tdata := t.int64s()
		odata := other.int64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = int64(math.Pow(float64(tdata[i]), float64(odata[j])))
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = int64(math.Pow(float64(tdata[i]), float64(odata[j])))
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = int64(math.Pow(float64(tdata[i]), float64(odata[j])))
				i++
			}
		default:
			vecPowI64(tdata, odata)
		}
	case reflect.Uint:
		tdata := t.uints()
		odata := other.uints()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = uint(math.Pow(float64(tdata[i]), float64(odata[j])))
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = uint(math.Pow(float64(tdata[i]), float64(odata[j])))
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = uint(math.Pow(float64(tdata[i]), float64(odata[j])))
				i++
			}
		default:
			vecPowU(tdata, odata)
		}
	case reflect.Uint8:
		tdata := t.uint8s()
		odata := other.uint8s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = uint8(math.Pow(float64(tdata[i]), float64(odata[j])))
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = uint8(math.Pow(float64(tdata[i]), float64(odata[j])))
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = uint8(math.Pow(float64(tdata[i]), float64(odata[j])))
				i++
			}
		default:
			vecPowU8(tdata, odata)
		}
	case reflect.Uint16:
		tdata := t.uint16s()
		odata := other.uint16s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = uint16(math.Pow(float64(tdata[i]), float64(odata[j])))
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = uint16(math.Pow(float64(tdata[i]), float64(odata[j])))
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = uint16(math.Pow(float64(tdata[i]), float64(odata[j])))
				i++
			}
		default:
			vecPowU16(tdata, odata)
		}
	case reflect.Uint32:
		tdata := t.uint32s()
		odata := other.uint32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = uint32(math.Pow(float64(tdata[i]), float64(odata[j])))
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = uint32(math.Pow(float64(tdata[i]), float64(odata[j])))
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = uint32(math.Pow(float64(tdata[i]), float64(odata[j])))
				i++
			}
		default:
			vecPowU32(tdata, odata)
		}
	case reflect.Uint64:
		tdata := t.uint64s()
		odata := other.uint64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = uint64(math.Pow(float64(tdata[i]), float64(odata[j])))
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = uint64(math.Pow(float64(tdata[i]), float64(odata[j])))
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = uint64(math.Pow(float64(tdata[i]), float64(odata[j])))
				i++
			}
		default:
			vecPowU64(tdata, odata)
		}
	case reflect.Float32:
		tdata := t.float32s()
		odata := other.float32s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = math32.Pow(tdata[i], odata[j])
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = math32.Pow(tdata[i], odata[j])
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = math32.Pow(tdata[i], odata[j])
				i++
			}
		default:
			vecPowF32(tdata, odata)
		}
	case reflect.Float64:
		tdata := t.float64s()
		odata := other.float64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = math.Pow(tdata[i], odata[j])
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = math.Pow(tdata[i], odata[j])
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = math.Pow(tdata[i], odata[j])
				i++
			}
		default:
			vecPowF64(tdata, odata)
		}
	case reflect.Complex64:
		tdata := t.complex64s()
		odata := other.complex64s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = complex64(cmplx.Pow(complex128(tdata[i]), complex128(odata[j])))
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = complex64(cmplx.Pow(complex128(tdata[i]), complex128(odata[j])))
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = complex64(cmplx.Pow(complex128(tdata[i]), complex128(odata[j])))
				i++
			}
		default:
			vecPowC64(tdata, odata)
		}
	case reflect.Complex128:
		tdata := t.complex128s()
		odata := other.complex128s()
		var i, j int
		switch {
		case it != nil && ot != nil:
			for {
				if i, err = it.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				if j, err = ot.Next(); err != nil {
					err = handleNoOp(err)
					break
				}
				tdata[i] = cmplx.Pow(tdata[i], odata[j])
			}
		case it != nil && ot == nil:
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				tdata[i] = cmplx.Pow(tdata[i], odata[j])
				j++
			}
			err = handleNoOp(err)
		case it == nil && ot != nil:
			for j, err = ot.Next(); err == nil; j, err = ot.Next() {
				tdata[i] = cmplx.Pow(tdata[i], odata[j])
				i++
			}
		default:
			vecPowC128(tdata, odata)
		}
	default:
		// TODO: Handle Number interface
	}

	return
}

/* Trans */

// Trans performs addition on a *Dense and a scalar value. The scalar value has to be of the same
// type as defined in the *Dense, otherwise an error will be returned.
func (t *Dense) Trans(other interface{}, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepUnaryDense(t, opts...)
	if err != nil {
		return nil, err
	}

	switch {
	case incr:
		switch t.t.Kind() {
		case reflect.Int:
			err = incrTransI(t.ints(), reuse.ints(), other.(int))
			retVal = reuse
		case reflect.Int8:
			err = incrTransI8(t.int8s(), reuse.int8s(), other.(int8))
			retVal = reuse
		case reflect.Int16:
			err = incrTransI16(t.int16s(), reuse.int16s(), other.(int16))
			retVal = reuse
		case reflect.Int32:
			err = incrTransI32(t.int32s(), reuse.int32s(), other.(int32))
			retVal = reuse
		case reflect.Int64:
			err = incrTransI64(t.int64s(), reuse.int64s(), other.(int64))
			retVal = reuse
		case reflect.Uint:
			err = incrTransU(t.uints(), reuse.uints(), other.(uint))
			retVal = reuse
		case reflect.Uint8:
			err = incrTransU8(t.uint8s(), reuse.uint8s(), other.(uint8))
			retVal = reuse
		case reflect.Uint16:
			err = incrTransU16(t.uint16s(), reuse.uint16s(), other.(uint16))
			retVal = reuse
		case reflect.Uint32:
			err = incrTransU32(t.uint32s(), reuse.uint32s(), other.(uint32))
			retVal = reuse
		case reflect.Uint64:
			err = incrTransU64(t.uint64s(), reuse.uint64s(), other.(uint64))
			retVal = reuse
		case reflect.Float32:
			err = incrTransF32(t.float32s(), reuse.float32s(), other.(float32))
			retVal = reuse
		case reflect.Float64:
			err = incrTransF64(t.float64s(), reuse.float64s(), other.(float64))
			retVal = reuse
		case reflect.Complex64:
			err = incrTransC64(t.complex64s(), reuse.complex64s(), other.(complex64))
			retVal = reuse
		case reflect.Complex128:
			err = incrTransC128(t.complex128s(), reuse.complex128s(), other.(complex128))
			retVal = reuse
		}
	case toReuse:
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t)
		}
		reuse.trans(other)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		retVal.trans(other)
	case !safe:
		t.trans(other)
		retVal = t
	}
	return
}
func (t *Dense) trans(other interface{}) (err error) {
	switch t.t.Kind() {
	case reflect.Int:
		b := other.(int)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.ints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transI(t.ints(), b)
	case reflect.Int8:
		b := other.(int8)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transI8(t.int8s(), b)
	case reflect.Int16:
		b := other.(int16)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transI16(t.int16s(), b)
	case reflect.Int32:
		b := other.(int32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transI32(t.int32s(), b)
	case reflect.Int64:
		b := other.(int64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transI64(t.int64s(), b)
	case reflect.Uint:
		b := other.(uint)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transU(t.uints(), b)
	case reflect.Uint8:
		b := other.(uint8)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transU8(t.uint8s(), b)
	case reflect.Uint16:
		b := other.(uint16)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transU16(t.uint16s(), b)
	case reflect.Uint32:
		b := other.(uint32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transU32(t.uint32s(), b)
	case reflect.Uint64:
		b := other.(uint64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transU64(t.uint64s(), b)
	case reflect.Float32:
		b := other.(float32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transF32(t.float32s(), b)
	case reflect.Float64:
		b := other.(float64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transF64(t.float64s(), b)
	case reflect.Complex64:
		b := other.(complex64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transC64(t.complex64s(), b)
	case reflect.Complex128:
		b := other.(complex128)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex128s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] + b
			}
			return nil
		}
		transC128(t.complex128s(), b)
	}
	return nil
}

/* TransInv */

// TransInv performs subtraction on a *Dense and a scalar value. The scalar value has to be of the same
// type as defined in the *Dense, otherwise an error will be returned.
func (t *Dense) TransInv(other interface{}, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepUnaryDense(t, opts...)
	if err != nil {
		return nil, err
	}

	switch {
	case incr:
		switch t.t.Kind() {
		case reflect.Int:
			err = incrTransInvI(t.ints(), reuse.ints(), other.(int))
			retVal = reuse
		case reflect.Int8:
			err = incrTransInvI8(t.int8s(), reuse.int8s(), other.(int8))
			retVal = reuse
		case reflect.Int16:
			err = incrTransInvI16(t.int16s(), reuse.int16s(), other.(int16))
			retVal = reuse
		case reflect.Int32:
			err = incrTransInvI32(t.int32s(), reuse.int32s(), other.(int32))
			retVal = reuse
		case reflect.Int64:
			err = incrTransInvI64(t.int64s(), reuse.int64s(), other.(int64))
			retVal = reuse
		case reflect.Uint:
			err = incrTransInvU(t.uints(), reuse.uints(), other.(uint))
			retVal = reuse
		case reflect.Uint8:
			err = incrTransInvU8(t.uint8s(), reuse.uint8s(), other.(uint8))
			retVal = reuse
		case reflect.Uint16:
			err = incrTransInvU16(t.uint16s(), reuse.uint16s(), other.(uint16))
			retVal = reuse
		case reflect.Uint32:
			err = incrTransInvU32(t.uint32s(), reuse.uint32s(), other.(uint32))
			retVal = reuse
		case reflect.Uint64:
			err = incrTransInvU64(t.uint64s(), reuse.uint64s(), other.(uint64))
			retVal = reuse
		case reflect.Float32:
			err = incrTransInvF32(t.float32s(), reuse.float32s(), other.(float32))
			retVal = reuse
		case reflect.Float64:
			err = incrTransInvF64(t.float64s(), reuse.float64s(), other.(float64))
			retVal = reuse
		case reflect.Complex64:
			err = incrTransInvC64(t.complex64s(), reuse.complex64s(), other.(complex64))
			retVal = reuse
		case reflect.Complex128:
			err = incrTransInvC128(t.complex128s(), reuse.complex128s(), other.(complex128))
			retVal = reuse
		}
	case toReuse:
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t)
		}
		reuse.transinv(other)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		retVal.transinv(other)
	case !safe:
		t.transinv(other)
		retVal = t
	}
	return
}
func (t *Dense) transinv(other interface{}) (err error) {
	switch t.t.Kind() {
	case reflect.Int:
		b := other.(int)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.ints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvI(t.ints(), b)
	case reflect.Int8:
		b := other.(int8)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvI8(t.int8s(), b)
	case reflect.Int16:
		b := other.(int16)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvI16(t.int16s(), b)
	case reflect.Int32:
		b := other.(int32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvI32(t.int32s(), b)
	case reflect.Int64:
		b := other.(int64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvI64(t.int64s(), b)
	case reflect.Uint:
		b := other.(uint)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvU(t.uints(), b)
	case reflect.Uint8:
		b := other.(uint8)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvU8(t.uint8s(), b)
	case reflect.Uint16:
		b := other.(uint16)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvU16(t.uint16s(), b)
	case reflect.Uint32:
		b := other.(uint32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvU32(t.uint32s(), b)
	case reflect.Uint64:
		b := other.(uint64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvU64(t.uint64s(), b)
	case reflect.Float32:
		b := other.(float32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvF32(t.float32s(), b)
	case reflect.Float64:
		b := other.(float64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvF64(t.float64s(), b)
	case reflect.Complex64:
		b := other.(complex64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvC64(t.complex64s(), b)
	case reflect.Complex128:
		b := other.(complex128)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex128s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvC128(t.complex128s(), b)
	}
	return nil
}

/* TransInvR */

// TransInvR performs subtraction on a *Dense and a scalar value. The scalar value has to be of the same
// type as defined in the *Dense, otherwise an error will be returned.
func (t *Dense) TransInvR(other interface{}, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepUnaryDense(t, opts...)
	if err != nil {
		return nil, err
	}

	switch {
	case incr:
		switch t.t.Kind() {
		case reflect.Int:
			err = incrTransInvRI(t.ints(), reuse.ints(), other.(int))
			retVal = reuse
		case reflect.Int8:
			err = incrTransInvRI8(t.int8s(), reuse.int8s(), other.(int8))
			retVal = reuse
		case reflect.Int16:
			err = incrTransInvRI16(t.int16s(), reuse.int16s(), other.(int16))
			retVal = reuse
		case reflect.Int32:
			err = incrTransInvRI32(t.int32s(), reuse.int32s(), other.(int32))
			retVal = reuse
		case reflect.Int64:
			err = incrTransInvRI64(t.int64s(), reuse.int64s(), other.(int64))
			retVal = reuse
		case reflect.Uint:
			err = incrTransInvRU(t.uints(), reuse.uints(), other.(uint))
			retVal = reuse
		case reflect.Uint8:
			err = incrTransInvRU8(t.uint8s(), reuse.uint8s(), other.(uint8))
			retVal = reuse
		case reflect.Uint16:
			err = incrTransInvRU16(t.uint16s(), reuse.uint16s(), other.(uint16))
			retVal = reuse
		case reflect.Uint32:
			err = incrTransInvRU32(t.uint32s(), reuse.uint32s(), other.(uint32))
			retVal = reuse
		case reflect.Uint64:
			err = incrTransInvRU64(t.uint64s(), reuse.uint64s(), other.(uint64))
			retVal = reuse
		case reflect.Float32:
			err = incrTransInvRF32(t.float32s(), reuse.float32s(), other.(float32))
			retVal = reuse
		case reflect.Float64:
			err = incrTransInvRF64(t.float64s(), reuse.float64s(), other.(float64))
			retVal = reuse
		case reflect.Complex64:
			err = incrTransInvRC64(t.complex64s(), reuse.complex64s(), other.(complex64))
			retVal = reuse
		case reflect.Complex128:
			err = incrTransInvRC128(t.complex128s(), reuse.complex128s(), other.(complex128))
			retVal = reuse
		}
	case toReuse:
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t)
		}
		reuse.transinvr(other)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		retVal.transinvr(other)
	case !safe:
		t.transinvr(other)
		retVal = t
	}
	return
}
func (t *Dense) transinvr(other interface{}) (err error) {
	switch t.t.Kind() {
	case reflect.Int:
		b := other.(int)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.ints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrI(t.ints(), b)
	case reflect.Int8:
		b := other.(int8)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrI8(t.int8s(), b)
	case reflect.Int16:
		b := other.(int16)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrI16(t.int16s(), b)
	case reflect.Int32:
		b := other.(int32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrI32(t.int32s(), b)
	case reflect.Int64:
		b := other.(int64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrI64(t.int64s(), b)
	case reflect.Uint:
		b := other.(uint)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrU(t.uints(), b)
	case reflect.Uint8:
		b := other.(uint8)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrU8(t.uint8s(), b)
	case reflect.Uint16:
		b := other.(uint16)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrU16(t.uint16s(), b)
	case reflect.Uint32:
		b := other.(uint32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrU32(t.uint32s(), b)
	case reflect.Uint64:
		b := other.(uint64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrU64(t.uint64s(), b)
	case reflect.Float32:
		b := other.(float32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrF32(t.float32s(), b)
	case reflect.Float64:
		b := other.(float64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrF64(t.float64s(), b)
	case reflect.Complex64:
		b := other.(complex64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrC64(t.complex64s(), b)
	case reflect.Complex128:
		b := other.(complex128)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex128s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] - b
			}
			return nil
		}
		transinvrC128(t.complex128s(), b)
	}
	return nil
}

/* Scale */

// Scale performs multiplication on a *Dense and a scalar value. The scalar value has to be of the same
// type as defined in the *Dense, otherwise an error will be returned.
func (t *Dense) Scale(other interface{}, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepUnaryDense(t, opts...)
	if err != nil {
		return nil, err
	}

	switch {
	case incr:
		switch t.t.Kind() {
		case reflect.Int:
			err = incrScaleI(t.ints(), reuse.ints(), other.(int))
			retVal = reuse
		case reflect.Int8:
			err = incrScaleI8(t.int8s(), reuse.int8s(), other.(int8))
			retVal = reuse
		case reflect.Int16:
			err = incrScaleI16(t.int16s(), reuse.int16s(), other.(int16))
			retVal = reuse
		case reflect.Int32:
			err = incrScaleI32(t.int32s(), reuse.int32s(), other.(int32))
			retVal = reuse
		case reflect.Int64:
			err = incrScaleI64(t.int64s(), reuse.int64s(), other.(int64))
			retVal = reuse
		case reflect.Uint:
			err = incrScaleU(t.uints(), reuse.uints(), other.(uint))
			retVal = reuse
		case reflect.Uint8:
			err = incrScaleU8(t.uint8s(), reuse.uint8s(), other.(uint8))
			retVal = reuse
		case reflect.Uint16:
			err = incrScaleU16(t.uint16s(), reuse.uint16s(), other.(uint16))
			retVal = reuse
		case reflect.Uint32:
			err = incrScaleU32(t.uint32s(), reuse.uint32s(), other.(uint32))
			retVal = reuse
		case reflect.Uint64:
			err = incrScaleU64(t.uint64s(), reuse.uint64s(), other.(uint64))
			retVal = reuse
		case reflect.Float32:
			err = incrScaleF32(t.float32s(), reuse.float32s(), other.(float32))
			retVal = reuse
		case reflect.Float64:
			err = incrScaleF64(t.float64s(), reuse.float64s(), other.(float64))
			retVal = reuse
		case reflect.Complex64:
			err = incrScaleC64(t.complex64s(), reuse.complex64s(), other.(complex64))
			retVal = reuse
		case reflect.Complex128:
			err = incrScaleC128(t.complex128s(), reuse.complex128s(), other.(complex128))
			retVal = reuse
		}
	case toReuse:
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t)
		}
		reuse.scale(other)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		retVal.scale(other)
	case !safe:
		t.scale(other)
		retVal = t
	}
	return
}
func (t *Dense) scale(other interface{}) (err error) {
	switch t.t.Kind() {
	case reflect.Int:
		b := other.(int)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.ints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleI(t.ints(), b)
	case reflect.Int8:
		b := other.(int8)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleI8(t.int8s(), b)
	case reflect.Int16:
		b := other.(int16)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleI16(t.int16s(), b)
	case reflect.Int32:
		b := other.(int32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleI32(t.int32s(), b)
	case reflect.Int64:
		b := other.(int64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleI64(t.int64s(), b)
	case reflect.Uint:
		b := other.(uint)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleU(t.uints(), b)
	case reflect.Uint8:
		b := other.(uint8)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleU8(t.uint8s(), b)
	case reflect.Uint16:
		b := other.(uint16)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleU16(t.uint16s(), b)
	case reflect.Uint32:
		b := other.(uint32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleU32(t.uint32s(), b)
	case reflect.Uint64:
		b := other.(uint64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleU64(t.uint64s(), b)
	case reflect.Float32:
		b := other.(float32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleF32(t.float32s(), b)
	case reflect.Float64:
		b := other.(float64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleF64(t.float64s(), b)
	case reflect.Complex64:
		b := other.(complex64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleC64(t.complex64s(), b)
	case reflect.Complex128:
		b := other.(complex128)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex128s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] * b
			}
			return nil
		}
		scaleC128(t.complex128s(), b)
	}
	return nil
}

/* ScaleInv */

// ScaleInv performs division on a *Dense and a scalar value. The scalar value has to be of the same
// type as defined in the *Dense, otherwise an error will be returned.
func (t *Dense) ScaleInv(other interface{}, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepUnaryDense(t, opts...)
	if err != nil {
		return nil, err
	}

	switch {
	case incr:
		switch t.t.Kind() {
		case reflect.Int:
			err = incrScaleInvI(t.ints(), reuse.ints(), other.(int))
			retVal = reuse
		case reflect.Int8:
			err = incrScaleInvI8(t.int8s(), reuse.int8s(), other.(int8))
			retVal = reuse
		case reflect.Int16:
			err = incrScaleInvI16(t.int16s(), reuse.int16s(), other.(int16))
			retVal = reuse
		case reflect.Int32:
			err = incrScaleInvI32(t.int32s(), reuse.int32s(), other.(int32))
			retVal = reuse
		case reflect.Int64:
			err = incrScaleInvI64(t.int64s(), reuse.int64s(), other.(int64))
			retVal = reuse
		case reflect.Uint:
			err = incrScaleInvU(t.uints(), reuse.uints(), other.(uint))
			retVal = reuse
		case reflect.Uint8:
			err = incrScaleInvU8(t.uint8s(), reuse.uint8s(), other.(uint8))
			retVal = reuse
		case reflect.Uint16:
			err = incrScaleInvU16(t.uint16s(), reuse.uint16s(), other.(uint16))
			retVal = reuse
		case reflect.Uint32:
			err = incrScaleInvU32(t.uint32s(), reuse.uint32s(), other.(uint32))
			retVal = reuse
		case reflect.Uint64:
			err = incrScaleInvU64(t.uint64s(), reuse.uint64s(), other.(uint64))
			retVal = reuse
		case reflect.Float32:
			err = incrScaleInvF32(t.float32s(), reuse.float32s(), other.(float32))
			retVal = reuse
		case reflect.Float64:
			err = incrScaleInvF64(t.float64s(), reuse.float64s(), other.(float64))
			retVal = reuse
		case reflect.Complex64:
			err = incrScaleInvC64(t.complex64s(), reuse.complex64s(), other.(complex64))
			retVal = reuse
		case reflect.Complex128:
			err = incrScaleInvC128(t.complex128s(), reuse.complex128s(), other.(complex128))
			retVal = reuse
		}
	case toReuse:
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t)
		}
		reuse.scaleinv(other)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		retVal.scaleinv(other)
	case !safe:
		t.scaleinv(other)
		retVal = t
	}
	return
}
func (t *Dense) scaleinv(other interface{}) (err error) {
	switch t.t.Kind() {
	case reflect.Int:
		b := other.(int)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.ints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvI(t.ints(), b)
	case reflect.Int8:
		b := other.(int8)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvI8(t.int8s(), b)
	case reflect.Int16:
		b := other.(int16)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvI16(t.int16s(), b)
	case reflect.Int32:
		b := other.(int32)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvI32(t.int32s(), b)
	case reflect.Int64:
		b := other.(int64)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvI64(t.int64s(), b)
	case reflect.Uint:
		b := other.(uint)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvU(t.uints(), b)
	case reflect.Uint8:
		b := other.(uint8)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvU8(t.uint8s(), b)
	case reflect.Uint16:
		b := other.(uint16)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvU16(t.uint16s(), b)
	case reflect.Uint32:
		b := other.(uint32)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvU32(t.uint32s(), b)
	case reflect.Uint64:
		b := other.(uint64)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvU64(t.uint64s(), b)
	case reflect.Float32:
		b := other.(float32)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvF32(t.float32s(), b)
	case reflect.Float64:
		b := other.(float64)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvF64(t.float64s(), b)
	case reflect.Complex64:
		b := other.(complex64)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvC64(t.complex64s(), b)
	case reflect.Complex128:
		b := other.(complex128)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex128s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvC128(t.complex128s(), b)
	}
	return nil
}

/* ScaleInvR */

// ScaleInvR performs division on a *Dense and a scalar value. The scalar value has to be of the same
// type as defined in the *Dense, otherwise an error will be returned.
func (t *Dense) ScaleInvR(other interface{}, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepUnaryDense(t, opts...)
	if err != nil {
		return nil, err
	}

	switch {
	case incr:
		switch t.t.Kind() {
		case reflect.Int:
			err = incrScaleInvRI(t.ints(), reuse.ints(), other.(int))
			retVal = reuse
		case reflect.Int8:
			err = incrScaleInvRI8(t.int8s(), reuse.int8s(), other.(int8))
			retVal = reuse
		case reflect.Int16:
			err = incrScaleInvRI16(t.int16s(), reuse.int16s(), other.(int16))
			retVal = reuse
		case reflect.Int32:
			err = incrScaleInvRI32(t.int32s(), reuse.int32s(), other.(int32))
			retVal = reuse
		case reflect.Int64:
			err = incrScaleInvRI64(t.int64s(), reuse.int64s(), other.(int64))
			retVal = reuse
		case reflect.Uint:
			err = incrScaleInvRU(t.uints(), reuse.uints(), other.(uint))
			retVal = reuse
		case reflect.Uint8:
			err = incrScaleInvRU8(t.uint8s(), reuse.uint8s(), other.(uint8))
			retVal = reuse
		case reflect.Uint16:
			err = incrScaleInvRU16(t.uint16s(), reuse.uint16s(), other.(uint16))
			retVal = reuse
		case reflect.Uint32:
			err = incrScaleInvRU32(t.uint32s(), reuse.uint32s(), other.(uint32))
			retVal = reuse
		case reflect.Uint64:
			err = incrScaleInvRU64(t.uint64s(), reuse.uint64s(), other.(uint64))
			retVal = reuse
		case reflect.Float32:
			err = incrScaleInvRF32(t.float32s(), reuse.float32s(), other.(float32))
			retVal = reuse
		case reflect.Float64:
			err = incrScaleInvRF64(t.float64s(), reuse.float64s(), other.(float64))
			retVal = reuse
		case reflect.Complex64:
			err = incrScaleInvRC64(t.complex64s(), reuse.complex64s(), other.(complex64))
			retVal = reuse
		case reflect.Complex128:
			err = incrScaleInvRC128(t.complex128s(), reuse.complex128s(), other.(complex128))
			retVal = reuse
		}
	case toReuse:
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t)
		}
		reuse.scaleinvr(other)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		retVal.scaleinvr(other)
	case !safe:
		t.scaleinvr(other)
		retVal = t
	}
	return
}
func (t *Dense) scaleinvr(other interface{}) (err error) {
	switch t.t.Kind() {
	case reflect.Int:
		b := other.(int)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.ints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrI(t.ints(), b)
	case reflect.Int8:
		b := other.(int8)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrI8(t.int8s(), b)
	case reflect.Int16:
		b := other.(int16)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrI16(t.int16s(), b)
	case reflect.Int32:
		b := other.(int32)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrI32(t.int32s(), b)
	case reflect.Int64:
		b := other.(int64)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrI64(t.int64s(), b)
	case reflect.Uint:
		b := other.(uint)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrU(t.uints(), b)
	case reflect.Uint8:
		b := other.(uint8)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrU8(t.uint8s(), b)
	case reflect.Uint16:
		b := other.(uint16)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrU16(t.uint16s(), b)
	case reflect.Uint32:
		b := other.(uint32)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrU32(t.uint32s(), b)
	case reflect.Uint64:
		b := other.(uint64)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrU64(t.uint64s(), b)
	case reflect.Float32:
		b := other.(float32)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrF32(t.float32s(), b)
	case reflect.Float64:
		b := other.(float64)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrF64(t.float64s(), b)
	case reflect.Complex64:
		b := other.(complex64)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrC64(t.complex64s(), b)
	case reflect.Complex128:
		b := other.(complex128)
		if t.IsMaterializable() {
			if b == 0 {
				err = t.zeroIter()
				if err != nil {
					err = errors.Wrapf(err, div0, -1)
					return
				}
				err = errors.Errorf(div0, -1)
				return
			}
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex128s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = data[i] / b
			}
			return nil
		}
		scaleinvrC128(t.complex128s(), b)
	}
	return nil
}

/* PowOf */

// PowOf performs exponentiation on a *Dense and a scalar value. The scalar value has to be of the same
// type as defined in the *Dense, otherwise an error will be returned.
func (t *Dense) PowOf(other interface{}, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepUnaryDense(t, opts...)
	if err != nil {
		return nil, err
	}

	switch {
	case incr:
		switch t.t.Kind() {
		case reflect.Int:
			err = incrPowOfI(t.ints(), reuse.ints(), other.(int))
			retVal = reuse
		case reflect.Int8:
			err = incrPowOfI8(t.int8s(), reuse.int8s(), other.(int8))
			retVal = reuse
		case reflect.Int16:
			err = incrPowOfI16(t.int16s(), reuse.int16s(), other.(int16))
			retVal = reuse
		case reflect.Int32:
			err = incrPowOfI32(t.int32s(), reuse.int32s(), other.(int32))
			retVal = reuse
		case reflect.Int64:
			err = incrPowOfI64(t.int64s(), reuse.int64s(), other.(int64))
			retVal = reuse
		case reflect.Uint:
			err = incrPowOfU(t.uints(), reuse.uints(), other.(uint))
			retVal = reuse
		case reflect.Uint8:
			err = incrPowOfU8(t.uint8s(), reuse.uint8s(), other.(uint8))
			retVal = reuse
		case reflect.Uint16:
			err = incrPowOfU16(t.uint16s(), reuse.uint16s(), other.(uint16))
			retVal = reuse
		case reflect.Uint32:
			err = incrPowOfU32(t.uint32s(), reuse.uint32s(), other.(uint32))
			retVal = reuse
		case reflect.Uint64:
			err = incrPowOfU64(t.uint64s(), reuse.uint64s(), other.(uint64))
			retVal = reuse
		case reflect.Float32:
			err = incrPowOfF32(t.float32s(), reuse.float32s(), other.(float32))
			retVal = reuse
		case reflect.Float64:
			err = incrPowOfF64(t.float64s(), reuse.float64s(), other.(float64))
			retVal = reuse
		case reflect.Complex64:
			err = incrPowOfC64(t.complex64s(), reuse.complex64s(), other.(complex64))
			retVal = reuse
		case reflect.Complex128:
			err = incrPowOfC128(t.complex128s(), reuse.complex128s(), other.(complex128))
			retVal = reuse
		}
	case toReuse:
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t)
		}
		reuse.powof(other)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		retVal.powof(other)
	case !safe:
		t.powof(other)
		retVal = t
	}
	return
}
func (t *Dense) powof(other interface{}) (err error) {
	switch t.t.Kind() {
	case reflect.Int:
		b := other.(int)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.ints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = int(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofI(t.ints(), b)
	case reflect.Int8:
		b := other.(int8)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = int8(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofI8(t.int8s(), b)
	case reflect.Int16:
		b := other.(int16)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = int16(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofI16(t.int16s(), b)
	case reflect.Int32:
		b := other.(int32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = int32(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofI32(t.int32s(), b)
	case reflect.Int64:
		b := other.(int64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = int64(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofI64(t.int64s(), b)
	case reflect.Uint:
		b := other.(uint)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = uint(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofU(t.uints(), b)
	case reflect.Uint8:
		b := other.(uint8)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = uint8(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofU8(t.uint8s(), b)
	case reflect.Uint16:
		b := other.(uint16)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = uint16(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofU16(t.uint16s(), b)
	case reflect.Uint32:
		b := other.(uint32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = uint32(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofU32(t.uint32s(), b)
	case reflect.Uint64:
		b := other.(uint64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = uint64(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofU64(t.uint64s(), b)
	case reflect.Float32:
		b := other.(float32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = math32.Pow(data[i], b)
			}
			return nil
		}
		powofF32(t.float32s(), b)
	case reflect.Float64:
		b := other.(float64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = math.Pow(data[i], b)
			}
			return nil
		}
		powofF64(t.float64s(), b)
	case reflect.Complex64:
		b := other.(complex64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = complex64(cmplx.Pow(complex128(data[i]), complex128(b)))
			}
			return nil
		}
		powofC64(t.complex64s(), b)
	case reflect.Complex128:
		b := other.(complex128)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex128s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = cmplx.Pow(data[i], b)
			}
			return nil
		}
		powofC128(t.complex128s(), b)
	}
	return nil
}

/* PowOfR */

// PowOfR performs exponentiation on a *Dense and a scalar value. The scalar value has to be of the same
// type as defined in the *Dense, otherwise an error will be returned.
func (t *Dense) PowOfR(other interface{}, opts ...FuncOpt) (retVal *Dense, err error) {
	reuse, safe, toReuse, incr, err := prepUnaryDense(t, opts...)
	if err != nil {
		return nil, err
	}

	switch {
	case incr:
		switch t.t.Kind() {
		case reflect.Int:
			err = incrPowOfRI(t.ints(), reuse.ints(), other.(int))
			retVal = reuse
		case reflect.Int8:
			err = incrPowOfRI8(t.int8s(), reuse.int8s(), other.(int8))
			retVal = reuse
		case reflect.Int16:
			err = incrPowOfRI16(t.int16s(), reuse.int16s(), other.(int16))
			retVal = reuse
		case reflect.Int32:
			err = incrPowOfRI32(t.int32s(), reuse.int32s(), other.(int32))
			retVal = reuse
		case reflect.Int64:
			err = incrPowOfRI64(t.int64s(), reuse.int64s(), other.(int64))
			retVal = reuse
		case reflect.Uint:
			err = incrPowOfRU(t.uints(), reuse.uints(), other.(uint))
			retVal = reuse
		case reflect.Uint8:
			err = incrPowOfRU8(t.uint8s(), reuse.uint8s(), other.(uint8))
			retVal = reuse
		case reflect.Uint16:
			err = incrPowOfRU16(t.uint16s(), reuse.uint16s(), other.(uint16))
			retVal = reuse
		case reflect.Uint32:
			err = incrPowOfRU32(t.uint32s(), reuse.uint32s(), other.(uint32))
			retVal = reuse
		case reflect.Uint64:
			err = incrPowOfRU64(t.uint64s(), reuse.uint64s(), other.(uint64))
			retVal = reuse
		case reflect.Float32:
			err = incrPowOfRF32(t.float32s(), reuse.float32s(), other.(float32))
			retVal = reuse
		case reflect.Float64:
			err = incrPowOfRF64(t.float64s(), reuse.float64s(), other.(float64))
			retVal = reuse
		case reflect.Complex64:
			err = incrPowOfRC64(t.complex64s(), reuse.complex64s(), other.(complex64))
			retVal = reuse
		case reflect.Complex128:
			err = incrPowOfRC128(t.complex128s(), reuse.complex128s(), other.(complex128))
			retVal = reuse
		}
	case toReuse:
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			copyDenseIter(reuse, t, nil, it)
		} else {
			copyDense(reuse, t)
		}
		reuse.powofr(other)
		retVal = reuse
	case safe:
		if t.IsMaterializable() {
			retVal = t.Materialize().(*Dense)
		} else {
			retVal = t.Clone().(*Dense)
		}
		retVal.powofr(other)
	case !safe:
		t.powofr(other)
		retVal = t
	}
	return
}
func (t *Dense) powofr(other interface{}) (err error) {
	switch t.t.Kind() {
	case reflect.Int:
		b := other.(int)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.ints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = int(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofrI(t.ints(), b)
	case reflect.Int8:
		b := other.(int8)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = int8(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofrI8(t.int8s(), b)
	case reflect.Int16:
		b := other.(int16)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = int16(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofrI16(t.int16s(), b)
	case reflect.Int32:
		b := other.(int32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = int32(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofrI32(t.int32s(), b)
	case reflect.Int64:
		b := other.(int64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.int64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = int64(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofrI64(t.int64s(), b)
	case reflect.Uint:
		b := other.(uint)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uints()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = uint(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofrU(t.uints(), b)
	case reflect.Uint8:
		b := other.(uint8)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint8s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = uint8(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofrU8(t.uint8s(), b)
	case reflect.Uint16:
		b := other.(uint16)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint16s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = uint16(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofrU16(t.uint16s(), b)
	case reflect.Uint32:
		b := other.(uint32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = uint32(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofrU32(t.uint32s(), b)
	case reflect.Uint64:
		b := other.(uint64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.uint64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = uint64(math.Pow(float64(data[i]), float64(b)))
			}
			return nil
		}
		powofrU64(t.uint64s(), b)
	case reflect.Float32:
		b := other.(float32)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float32s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = math32.Pow(data[i], b)
			}
			return nil
		}
		powofrF32(t.float32s(), b)
	case reflect.Float64:
		b := other.(float64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.float64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = math.Pow(data[i], b)
			}
			return nil
		}
		powofrF64(t.float64s(), b)
	case reflect.Complex64:
		b := other.(complex64)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex64s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = complex64(cmplx.Pow(complex128(data[i]), complex128(b)))
			}
			return nil
		}
		powofrC64(t.complex64s(), b)
	case reflect.Complex128:
		b := other.(complex128)
		if t.IsMaterializable() {
			it := NewFlatIterator(t.AP)
			var i int
			data := t.complex128s()
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				data[i] = cmplx.Pow(data[i], b)
			}
			return nil
		}
		powofrC128(t.complex128s(), b)
	}
	return nil
}
