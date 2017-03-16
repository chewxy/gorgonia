package tensor

import (
	"reflect"
	"testing"
	"testing/quick"
)

/*
GENERATED FILE. DO NOT EDIT
*/

/* Eq */

func TestDense_eqDD_Transitivity(t *testing.T) {
	fB := func(a, b, c *QCDenseB) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for bool failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for bool failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for bool failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fB, nil); err != nil {
		t.Error(err)
	}

	fIterB := func(a, b, c *QCDenseB) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// b iter
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// both a and b iter
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterB, nil); err != nil {
		t.Error(err)
	}
	fI := func(a, b, c *QCDenseI) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}

	fIterI := func(a, b, c *QCDenseI) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, b, c *QCDenseI8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}

	fIterI8 := func(a, b, c *QCDenseI8) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, b, c *QCDenseI16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}

	fIterI16 := func(a, b, c *QCDenseI16) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, b, c *QCDenseI32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}

	fIterI32 := func(a, b, c *QCDenseI32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, b, c *QCDenseI64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}

	fIterI64 := func(a, b, c *QCDenseI64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, b, c *QCDenseU) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}

	fIterU := func(a, b, c *QCDenseU) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, b, c *QCDenseU8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}

	fIterU8 := func(a, b, c *QCDenseU8) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, b, c *QCDenseU16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}

	fIterU16 := func(a, b, c *QCDenseU16) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, b, c *QCDenseU32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}

	fIterU32 := func(a, b, c *QCDenseU32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, b, c *QCDenseU64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}

	fIterU64 := func(a, b, c *QCDenseU64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU64, nil); err != nil {
		t.Error(err)
	}
	fUintptr := func(a, b, c *QCDenseUintptr) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uintptr failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uintptr failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uintptr failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fUintptr, nil); err != nil {
		t.Error(err)
	}

	fIterUintptr := func(a, b, c *QCDenseUintptr) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// b iter
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// both a and b iter
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterUintptr, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, b, c *QCDenseF32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for float32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for float32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for float32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}

	fIterF32 := func(a, b, c *QCDenseF32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, b, c *QCDenseF64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for float64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for float64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for float64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}

	fIterF64 := func(a, b, c *QCDenseF64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF64, nil); err != nil {
		t.Error(err)
	}
	fC64 := func(a, b, c *QCDenseC64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for complex64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for complex64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for complex64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fC64, nil); err != nil {
		t.Error(err)
	}

	fIterC64 := func(a, b, c *QCDenseC64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []complex64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.complex64s()
		bcs = bxc.complex64s()
		acs = axc.complex64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.complex64s()
		bcs = bxc.complex64s()
		acs = axc.complex64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.complex64s()
		bcs = bxc.complex64s()
		acs = axc.complex64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterC64, nil); err != nil {
		t.Error(err)
	}
	fC128 := func(a, b, c *QCDenseC128) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for complex128 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for complex128 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for complex128 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fC128, nil); err != nil {
		t.Error(err)
	}

	fIterC128 := func(a, b, c *QCDenseC128) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []complex128
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b2, AsSameType())
		bxc, _ = b1.eqDD(c2, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.complex128s()
		bcs = bxc.complex128s()
		acs = axc.complex128s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.eqDD(b1, AsSameType())
		bxc, _ = b2.eqDD(c1, AsSameType())
		axc, _ = a2.eqDD(c1, AsSameType())

		abs = axb.complex128s()
		bcs = bxc.complex128s()
		acs = axc.complex128s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDD(b1, AsSameType())
		bxc, _ = b1.eqDD(c1, AsSameType())
		axc, _ = a1.eqDD(c1, AsSameType())

		abs = axb.complex128s()
		bcs = bxc.complex128s()
		acs = axc.complex128s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterC128, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, b, c *QCDenseStr) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for string failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for string failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for string failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}

	fIterStr := func(a, b, c *QCDenseStr) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// b iter
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// both a and b iter
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterStr, nil); err != nil {
		t.Error(err)
	}
	fUnsafePointer := func(a, b, c *QCDenseUnsafePointer) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq transitivity for unsafe.Pointer failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for unsafe.Pointer failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for unsafe.Pointer failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fUnsafePointer, nil); err != nil {
		t.Error(err)
	}

	fIterUnsafePointer := func(a, b, c *QCDenseUnsafePointer) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.eqDD(b2)
		bxc, _ = b1.eqDD(c2)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// b iter
		axb, _ = a2.eqDD(b1)
		bxc, _ = b2.eqDD(c1)
		axc, _ = a2.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// both a and b iter
		axb, _ = a1.eqDD(b1)
		bxc, _ = b1.eqDD(c1)
		axc, _ = a1.eqDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterUnsafePointer, nil); err != nil {
		t.Error(err)
	}
}

func Test_Dense_eqDD_funcOpts(t *testing.T) {
	fB := func(a, b *QCDenseB) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for bool failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for bool failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fB, nil); err != nil {
		t.Error(err)
	}
	fI := func(a, b *QCDenseI) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for int failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for int failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for int failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for int failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for int failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, b *QCDenseI8) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for int8 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for int8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int8, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for int8 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for int8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for int8 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, b *QCDenseI16) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for int16 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for int16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int16, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for int16 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for int16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for int16 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, b *QCDenseI32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for int32 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for int32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int32, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for int32 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for int32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for int32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, b *QCDenseI64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for int64 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for int64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int64, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for int64 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for int64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for int64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, b *QCDenseU) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for uint failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for uint failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for uint failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for uint failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for uint failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, b *QCDenseU8) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for uint8 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for uint8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint8, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for uint8 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for uint8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for uint8 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, b *QCDenseU16) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for uint16 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for uint16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint16, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for uint16 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for uint16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for uint16 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, b *QCDenseU32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for uint32 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for uint32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint32, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for uint32 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for uint32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for uint32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, b *QCDenseU64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for uint64 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for uint64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint64, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for uint64 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for uint64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for uint64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}
	fUintptr := func(a, b *QCDenseUintptr) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for uintptr failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for uintptr failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fUintptr, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, b *QCDenseF32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for float32 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for float32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Float32, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for float32 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for float32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for float32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, b *QCDenseF64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for float64 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for float64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Float64, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for float64 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for float64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for float64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}
	fC64 := func(a, b *QCDenseC64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for complex64 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for complex64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Complex64, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for complex64 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for complex64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for complex64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fC64, nil); err != nil {
		t.Error(err)
	}
	fC128 := func(a, b *QCDenseC128) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for complex128 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for complex128 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Complex128, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Eq as same type reuse for complex128 failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq as same type reuse for complex128 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.eqDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Eq for complex128 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fC128, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, b *QCDenseStr) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for string failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for string failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}
	fUnsafePointer := func(a, b *QCDenseUnsafePointer) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.eqDD(b.Dense); err != nil {
			t.Errorf("Test Eq reuse for unsafe.Pointer failed(axb): %v", err)
		}
		if ret, err = a.eqDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Eq reuse for unsafe.Pointer failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fUnsafePointer, nil); err != nil {
		t.Error(err)
	}
}

/* Ne */

func TestDense_NeDD_Symmetry(t *testing.T) {
	fB := func(a, b *QCDenseB) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for bool failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for bool failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fB, nil); err != nil {
		t.Error(err)
	}
	fIterB := func(a, b, c *QCDenseB) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// b iter
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// both a and b iter
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		return true
	}
	if err := quick.Check(fIterB, nil); err != nil {
		t.Error(err)
	}
	fI := func(a, b *QCDenseI) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for int failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for int failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}
	fIterI := func(a, b, c *QCDenseI) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []int
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.ints()
		bas = bxa.ints()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.ints()
		bas = bxa.ints()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.ints()
		bas = bxa.ints()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, b *QCDenseI8) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for int8 failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for int8 failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}
	fIterI8 := func(a, b, c *QCDenseI8) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []int8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.int8s()
		bas = bxa.int8s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.int8s()
		bas = bxa.int8s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.int8s()
		bas = bxa.int8s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, b *QCDenseI16) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for int16 failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for int16 failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}
	fIterI16 := func(a, b, c *QCDenseI16) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []int16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.int16s()
		bas = bxa.int16s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.int16s()
		bas = bxa.int16s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.int16s()
		bas = bxa.int16s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, b *QCDenseI32) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for int32 failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for int32 failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}
	fIterI32 := func(a, b, c *QCDenseI32) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []int32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.int32s()
		bas = bxa.int32s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.int32s()
		bas = bxa.int32s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.int32s()
		bas = bxa.int32s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, b *QCDenseI64) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for int64 failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for int64 failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}
	fIterI64 := func(a, b, c *QCDenseI64) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []int64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.int64s()
		bas = bxa.int64s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.int64s()
		bas = bxa.int64s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.int64s()
		bas = bxa.int64s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, b *QCDenseU) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for uint failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for uint failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}
	fIterU := func(a, b, c *QCDenseU) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []uint
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.uints()
		bas = bxa.uints()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.uints()
		bas = bxa.uints()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.uints()
		bas = bxa.uints()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, b *QCDenseU8) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for uint8 failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for uint8 failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}
	fIterU8 := func(a, b, c *QCDenseU8) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []uint8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.uint8s()
		bas = bxa.uint8s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.uint8s()
		bas = bxa.uint8s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.uint8s()
		bas = bxa.uint8s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, b *QCDenseU16) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for uint16 failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for uint16 failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}
	fIterU16 := func(a, b, c *QCDenseU16) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []uint16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.uint16s()
		bas = bxa.uint16s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.uint16s()
		bas = bxa.uint16s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.uint16s()
		bas = bxa.uint16s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, b *QCDenseU32) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for uint32 failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for uint32 failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}
	fIterU32 := func(a, b, c *QCDenseU32) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []uint32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.uint32s()
		bas = bxa.uint32s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.uint32s()
		bas = bxa.uint32s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.uint32s()
		bas = bxa.uint32s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, b *QCDenseU64) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for uint64 failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for uint64 failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}
	fIterU64 := func(a, b, c *QCDenseU64) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []uint64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.uint64s()
		bas = bxa.uint64s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.uint64s()
		bas = bxa.uint64s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.uint64s()
		bas = bxa.uint64s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterU64, nil); err != nil {
		t.Error(err)
	}
	fUintptr := func(a, b *QCDenseUintptr) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for uintptr failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for uintptr failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fUintptr, nil); err != nil {
		t.Error(err)
	}
	fIterUintptr := func(a, b, c *QCDenseUintptr) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// b iter
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// both a and b iter
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		return true
	}
	if err := quick.Check(fIterUintptr, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, b *QCDenseF32) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for float32 failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for float32 failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}
	fIterF32 := func(a, b, c *QCDenseF32) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []float32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.float32s()
		bas = bxa.float32s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.float32s()
		bas = bxa.float32s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.float32s()
		bas = bxa.float32s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, b *QCDenseF64) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for float64 failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for float64 failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}
	fIterF64 := func(a, b, c *QCDenseF64) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []float64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.float64s()
		bas = bxa.float64s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.float64s()
		bas = bxa.float64s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.float64s()
		bas = bxa.float64s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterF64, nil); err != nil {
		t.Error(err)
	}
	fC64 := func(a, b *QCDenseC64) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for complex64 failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for complex64 failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fC64, nil); err != nil {
		t.Error(err)
	}
	fIterC64 := func(a, b, c *QCDenseC64) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []complex64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.complex64s()
		bas = bxa.complex64s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.complex64s()
		bas = bxa.complex64s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.complex64s()
		bas = bxa.complex64s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterC64, nil); err != nil {
		t.Error(err)
	}
	fC128 := func(a, b *QCDenseC128) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for complex128 failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for complex128 failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fC128, nil); err != nil {
		t.Error(err)
	}
	fIterC128 := func(a, b, c *QCDenseC128) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		var abs, bas []complex128
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b2, AsSameType())
		bxa, _ = b1.neDD(a2, AsSameType())

		abs = axb.complex128s()
		bas = bxa.complex128s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// b iter bools
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a2.neDD(b1, AsSameType())
		bxa, _ = b2.neDD(a1, AsSameType())

		abs = axb.complex128s()
		bas = bxa.complex128s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		// both a and b iter bools
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// a iter asSame
		axb, _ = a1.neDD(b1, AsSameType())
		bxa, _ = b1.neDD(a1, AsSameType())

		abs = axb.complex128s()
		bas = bxa.complex128s()

		for i, vab := range abs {
			if vab != bas[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fIterC128, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, b *QCDenseStr) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for string failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for string failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}
	fIterStr := func(a, b, c *QCDenseStr) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// b iter
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// both a and b iter
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		return true
	}
	if err := quick.Check(fIterStr, nil); err != nil {
		t.Error(err)
	}
	fUnsafePointer := func(a, b *QCDenseUnsafePointer) bool {
		var axb, bxa *Dense
		var err error
		if axb, err = a.neDD(b.Dense); err != nil {
			t.Errorf("Test Ne transitivity for unsafe.Pointer failed (axb) : %v ", err)
			return false
		}
		if bxa, err = b.neDD(a.Dense); err != nil {
			t.Errorf("Test Ne transitivity for unsafe.Pointer failed (bxa): %v", err)
			return false
		}

		ab := axb.bools()
		ba := bxa.bools()

		for i, vab := range ab {
			if vab != ba[i] {
				return false
			}
		}
		return true
	}
	if err := quick.Check(fUnsafePointer, nil); err != nil {
		t.Error(err)
	}
	fIterUnsafePointer := func(a, b, c *QCDenseUnsafePointer) bool {
		var axb, bxa *Dense
		var a1, b1 *Dense // sliced
		var a2, b2 *Dense // materialized slice

		var abb, bab []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.neDD(b2)
		bxa, _ = b1.neDD(a2)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// b iter
		axb, _ = a2.neDD(b1)
		bxa, _ = b2.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		// both a and b iter
		axb, _ = a1.neDD(b1)
		bxa, _ = b1.neDD(a1)

		abb = axb.bools()
		bab = bxa.bools()

		for i, vab := range abb {
			if vab != bab[i] {
				return false
			}
		}

		return true
	}
	if err := quick.Check(fIterUnsafePointer, nil); err != nil {
		t.Error(err)
	}
}

/* Gt */

func TestDense_gtDD_Transitivity(t *testing.T) {
	fI := func(a, b, c *QCDenseI) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}

	fIterI := func(a, b, c *QCDenseI) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b2, AsSameType())
		bxc, _ = b1.gtDD(c2, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gtDD(b1, AsSameType())
		bxc, _ = b2.gtDD(c1, AsSameType())
		axc, _ = a2.gtDD(c1, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b1, AsSameType())
		bxc, _ = b1.gtDD(c1, AsSameType())
		axc, _ = a1.gtDD(c1, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, b, c *QCDenseI8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}

	fIterI8 := func(a, b, c *QCDenseI8) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b2, AsSameType())
		bxc, _ = b1.gtDD(c2, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gtDD(b1, AsSameType())
		bxc, _ = b2.gtDD(c1, AsSameType())
		axc, _ = a2.gtDD(c1, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b1, AsSameType())
		bxc, _ = b1.gtDD(c1, AsSameType())
		axc, _ = a1.gtDD(c1, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, b, c *QCDenseI16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}

	fIterI16 := func(a, b, c *QCDenseI16) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b2, AsSameType())
		bxc, _ = b1.gtDD(c2, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gtDD(b1, AsSameType())
		bxc, _ = b2.gtDD(c1, AsSameType())
		axc, _ = a2.gtDD(c1, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b1, AsSameType())
		bxc, _ = b1.gtDD(c1, AsSameType())
		axc, _ = a1.gtDD(c1, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, b, c *QCDenseI32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}

	fIterI32 := func(a, b, c *QCDenseI32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b2, AsSameType())
		bxc, _ = b1.gtDD(c2, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gtDD(b1, AsSameType())
		bxc, _ = b2.gtDD(c1, AsSameType())
		axc, _ = a2.gtDD(c1, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b1, AsSameType())
		bxc, _ = b1.gtDD(c1, AsSameType())
		axc, _ = a1.gtDD(c1, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, b, c *QCDenseI64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}

	fIterI64 := func(a, b, c *QCDenseI64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b2, AsSameType())
		bxc, _ = b1.gtDD(c2, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gtDD(b1, AsSameType())
		bxc, _ = b2.gtDD(c1, AsSameType())
		axc, _ = a2.gtDD(c1, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b1, AsSameType())
		bxc, _ = b1.gtDD(c1, AsSameType())
		axc, _ = a1.gtDD(c1, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, b, c *QCDenseU) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}

	fIterU := func(a, b, c *QCDenseU) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b2, AsSameType())
		bxc, _ = b1.gtDD(c2, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gtDD(b1, AsSameType())
		bxc, _ = b2.gtDD(c1, AsSameType())
		axc, _ = a2.gtDD(c1, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b1, AsSameType())
		bxc, _ = b1.gtDD(c1, AsSameType())
		axc, _ = a1.gtDD(c1, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, b, c *QCDenseU8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}

	fIterU8 := func(a, b, c *QCDenseU8) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b2, AsSameType())
		bxc, _ = b1.gtDD(c2, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gtDD(b1, AsSameType())
		bxc, _ = b2.gtDD(c1, AsSameType())
		axc, _ = a2.gtDD(c1, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b1, AsSameType())
		bxc, _ = b1.gtDD(c1, AsSameType())
		axc, _ = a1.gtDD(c1, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, b, c *QCDenseU16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}

	fIterU16 := func(a, b, c *QCDenseU16) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b2, AsSameType())
		bxc, _ = b1.gtDD(c2, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gtDD(b1, AsSameType())
		bxc, _ = b2.gtDD(c1, AsSameType())
		axc, _ = a2.gtDD(c1, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b1, AsSameType())
		bxc, _ = b1.gtDD(c1, AsSameType())
		axc, _ = a1.gtDD(c1, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, b, c *QCDenseU32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}

	fIterU32 := func(a, b, c *QCDenseU32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b2, AsSameType())
		bxc, _ = b1.gtDD(c2, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gtDD(b1, AsSameType())
		bxc, _ = b2.gtDD(c1, AsSameType())
		axc, _ = a2.gtDD(c1, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b1, AsSameType())
		bxc, _ = b1.gtDD(c1, AsSameType())
		axc, _ = a1.gtDD(c1, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, b, c *QCDenseU64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}

	fIterU64 := func(a, b, c *QCDenseU64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b2, AsSameType())
		bxc, _ = b1.gtDD(c2, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gtDD(b1, AsSameType())
		bxc, _ = b2.gtDD(c1, AsSameType())
		axc, _ = a2.gtDD(c1, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b1, AsSameType())
		bxc, _ = b1.gtDD(c1, AsSameType())
		axc, _ = a1.gtDD(c1, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU64, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, b, c *QCDenseF32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for float32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for float32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for float32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}

	fIterF32 := func(a, b, c *QCDenseF32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b2, AsSameType())
		bxc, _ = b1.gtDD(c2, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gtDD(b1, AsSameType())
		bxc, _ = b2.gtDD(c1, AsSameType())
		axc, _ = a2.gtDD(c1, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b1, AsSameType())
		bxc, _ = b1.gtDD(c1, AsSameType())
		axc, _ = a1.gtDD(c1, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, b, c *QCDenseF64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for float64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for float64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for float64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}

	fIterF64 := func(a, b, c *QCDenseF64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b2, AsSameType())
		bxc, _ = b1.gtDD(c2, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gtDD(b1, AsSameType())
		bxc, _ = b2.gtDD(c1, AsSameType())
		axc, _ = a2.gtDD(c1, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDD(b1, AsSameType())
		bxc, _ = b1.gtDD(c1, AsSameType())
		axc, _ = a1.gtDD(c1, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF64, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, b, c *QCDenseStr) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt transitivity for string failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for string failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for string failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}

	fIterStr := func(a, b, c *QCDenseStr) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.gtDD(b2)
		bxc, _ = b1.gtDD(c2)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// b iter
		axb, _ = a2.gtDD(b1)
		bxc, _ = b2.gtDD(c1)
		axc, _ = a2.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// both a and b iter
		axb, _ = a1.gtDD(b1)
		bxc, _ = b1.gtDD(c1)
		axc, _ = a1.gtDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterStr, nil); err != nil {
		t.Error(err)
	}
}

func Test_Dense_gtDD_funcOpts(t *testing.T) {
	fI := func(a, b *QCDenseI) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for int failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for int failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gt as same type reuse for int failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt as same type reuse for int failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gtDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gt for int failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, b *QCDenseI8) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for int8 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for int8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int8, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gt as same type reuse for int8 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt as same type reuse for int8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gtDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gt for int8 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, b *QCDenseI16) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for int16 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for int16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int16, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gt as same type reuse for int16 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt as same type reuse for int16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gtDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gt for int16 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, b *QCDenseI32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for int32 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for int32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int32, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gt as same type reuse for int32 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt as same type reuse for int32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gtDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gt for int32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, b *QCDenseI64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for int64 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for int64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int64, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gt as same type reuse for int64 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt as same type reuse for int64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gtDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gt for int64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, b *QCDenseU) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for uint failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for uint failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gt as same type reuse for uint failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt as same type reuse for uint failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gtDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gt for uint failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, b *QCDenseU8) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for uint8 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for uint8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint8, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gt as same type reuse for uint8 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt as same type reuse for uint8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gtDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gt for uint8 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, b *QCDenseU16) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for uint16 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for uint16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint16, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gt as same type reuse for uint16 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt as same type reuse for uint16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gtDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gt for uint16 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, b *QCDenseU32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for uint32 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for uint32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint32, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gt as same type reuse for uint32 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt as same type reuse for uint32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gtDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gt for uint32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, b *QCDenseU64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for uint64 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for uint64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint64, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gt as same type reuse for uint64 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt as same type reuse for uint64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gtDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gt for uint64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, b *QCDenseF32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for float32 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for float32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Float32, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gt as same type reuse for float32 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt as same type reuse for float32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gtDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gt for float32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, b *QCDenseF64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for float64 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for float64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Float64, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gt as same type reuse for float64 failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt as same type reuse for float64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gtDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gt for float64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, b *QCDenseStr) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gtDD(b.Dense); err != nil {
			t.Errorf("Test Gt reuse for string failed(axb): %v", err)
		}
		if ret, err = a.gtDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gt reuse for string failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}
}

/* Gte */

func TestDense_gteDD_Transitivity(t *testing.T) {
	fI := func(a, b, c *QCDenseI) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}

	fIterI := func(a, b, c *QCDenseI) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b2, AsSameType())
		bxc, _ = b1.gteDD(c2, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gteDD(b1, AsSameType())
		bxc, _ = b2.gteDD(c1, AsSameType())
		axc, _ = a2.gteDD(c1, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b1, AsSameType())
		bxc, _ = b1.gteDD(c1, AsSameType())
		axc, _ = a1.gteDD(c1, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, b, c *QCDenseI8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}

	fIterI8 := func(a, b, c *QCDenseI8) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b2, AsSameType())
		bxc, _ = b1.gteDD(c2, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gteDD(b1, AsSameType())
		bxc, _ = b2.gteDD(c1, AsSameType())
		axc, _ = a2.gteDD(c1, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b1, AsSameType())
		bxc, _ = b1.gteDD(c1, AsSameType())
		axc, _ = a1.gteDD(c1, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, b, c *QCDenseI16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}

	fIterI16 := func(a, b, c *QCDenseI16) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b2, AsSameType())
		bxc, _ = b1.gteDD(c2, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gteDD(b1, AsSameType())
		bxc, _ = b2.gteDD(c1, AsSameType())
		axc, _ = a2.gteDD(c1, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b1, AsSameType())
		bxc, _ = b1.gteDD(c1, AsSameType())
		axc, _ = a1.gteDD(c1, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, b, c *QCDenseI32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}

	fIterI32 := func(a, b, c *QCDenseI32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b2, AsSameType())
		bxc, _ = b1.gteDD(c2, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gteDD(b1, AsSameType())
		bxc, _ = b2.gteDD(c1, AsSameType())
		axc, _ = a2.gteDD(c1, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b1, AsSameType())
		bxc, _ = b1.gteDD(c1, AsSameType())
		axc, _ = a1.gteDD(c1, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, b, c *QCDenseI64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}

	fIterI64 := func(a, b, c *QCDenseI64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b2, AsSameType())
		bxc, _ = b1.gteDD(c2, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gteDD(b1, AsSameType())
		bxc, _ = b2.gteDD(c1, AsSameType())
		axc, _ = a2.gteDD(c1, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b1, AsSameType())
		bxc, _ = b1.gteDD(c1, AsSameType())
		axc, _ = a1.gteDD(c1, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, b, c *QCDenseU) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}

	fIterU := func(a, b, c *QCDenseU) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b2, AsSameType())
		bxc, _ = b1.gteDD(c2, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gteDD(b1, AsSameType())
		bxc, _ = b2.gteDD(c1, AsSameType())
		axc, _ = a2.gteDD(c1, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b1, AsSameType())
		bxc, _ = b1.gteDD(c1, AsSameType())
		axc, _ = a1.gteDD(c1, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, b, c *QCDenseU8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}

	fIterU8 := func(a, b, c *QCDenseU8) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b2, AsSameType())
		bxc, _ = b1.gteDD(c2, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gteDD(b1, AsSameType())
		bxc, _ = b2.gteDD(c1, AsSameType())
		axc, _ = a2.gteDD(c1, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b1, AsSameType())
		bxc, _ = b1.gteDD(c1, AsSameType())
		axc, _ = a1.gteDD(c1, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, b, c *QCDenseU16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}

	fIterU16 := func(a, b, c *QCDenseU16) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b2, AsSameType())
		bxc, _ = b1.gteDD(c2, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gteDD(b1, AsSameType())
		bxc, _ = b2.gteDD(c1, AsSameType())
		axc, _ = a2.gteDD(c1, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b1, AsSameType())
		bxc, _ = b1.gteDD(c1, AsSameType())
		axc, _ = a1.gteDD(c1, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, b, c *QCDenseU32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}

	fIterU32 := func(a, b, c *QCDenseU32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b2, AsSameType())
		bxc, _ = b1.gteDD(c2, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gteDD(b1, AsSameType())
		bxc, _ = b2.gteDD(c1, AsSameType())
		axc, _ = a2.gteDD(c1, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b1, AsSameType())
		bxc, _ = b1.gteDD(c1, AsSameType())
		axc, _ = a1.gteDD(c1, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, b, c *QCDenseU64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}

	fIterU64 := func(a, b, c *QCDenseU64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b2, AsSameType())
		bxc, _ = b1.gteDD(c2, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gteDD(b1, AsSameType())
		bxc, _ = b2.gteDD(c1, AsSameType())
		axc, _ = a2.gteDD(c1, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b1, AsSameType())
		bxc, _ = b1.gteDD(c1, AsSameType())
		axc, _ = a1.gteDD(c1, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU64, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, b, c *QCDenseF32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for float32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for float32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for float32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}

	fIterF32 := func(a, b, c *QCDenseF32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b2, AsSameType())
		bxc, _ = b1.gteDD(c2, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gteDD(b1, AsSameType())
		bxc, _ = b2.gteDD(c1, AsSameType())
		axc, _ = a2.gteDD(c1, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b1, AsSameType())
		bxc, _ = b1.gteDD(c1, AsSameType())
		axc, _ = a1.gteDD(c1, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, b, c *QCDenseF64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for float64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for float64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for float64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}

	fIterF64 := func(a, b, c *QCDenseF64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b2, AsSameType())
		bxc, _ = b1.gteDD(c2, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.gteDD(b1, AsSameType())
		bxc, _ = b2.gteDD(c1, AsSameType())
		axc, _ = a2.gteDD(c1, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDD(b1, AsSameType())
		bxc, _ = b1.gteDD(c1, AsSameType())
		axc, _ = a1.gteDD(c1, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF64, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, b, c *QCDenseStr) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte transitivity for string failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for string failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for string failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}

	fIterStr := func(a, b, c *QCDenseStr) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.gteDD(b2)
		bxc, _ = b1.gteDD(c2)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// b iter
		axb, _ = a2.gteDD(b1)
		bxc, _ = b2.gteDD(c1)
		axc, _ = a2.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// both a and b iter
		axb, _ = a1.gteDD(b1)
		bxc, _ = b1.gteDD(c1)
		axc, _ = a1.gteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterStr, nil); err != nil {
		t.Error(err)
	}
}

func Test_Dense_gteDD_funcOpts(t *testing.T) {
	fI := func(a, b *QCDenseI) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for int failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for int failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gte as same type reuse for int failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte as same type reuse for int failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gte for int failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, b *QCDenseI8) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for int8 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for int8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int8, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gte as same type reuse for int8 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte as same type reuse for int8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gte for int8 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, b *QCDenseI16) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for int16 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for int16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int16, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gte as same type reuse for int16 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte as same type reuse for int16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gte for int16 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, b *QCDenseI32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for int32 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for int32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int32, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gte as same type reuse for int32 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte as same type reuse for int32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gte for int32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, b *QCDenseI64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for int64 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for int64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int64, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gte as same type reuse for int64 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte as same type reuse for int64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gte for int64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, b *QCDenseU) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for uint failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for uint failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gte as same type reuse for uint failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte as same type reuse for uint failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gte for uint failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, b *QCDenseU8) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for uint8 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for uint8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint8, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gte as same type reuse for uint8 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte as same type reuse for uint8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gte for uint8 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, b *QCDenseU16) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for uint16 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for uint16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint16, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gte as same type reuse for uint16 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte as same type reuse for uint16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gte for uint16 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, b *QCDenseU32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for uint32 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for uint32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint32, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gte as same type reuse for uint32 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte as same type reuse for uint32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gte for uint32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, b *QCDenseU64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for uint64 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for uint64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint64, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gte as same type reuse for uint64 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte as same type reuse for uint64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gte for uint64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, b *QCDenseF32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for float32 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for float32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Float32, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gte as same type reuse for float32 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte as same type reuse for float32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gte for float32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, b *QCDenseF64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for float64 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for float64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Float64, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Gte as same type reuse for float64 failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte as same type reuse for float64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.gteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Gte for float64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, b *QCDenseStr) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.gteDD(b.Dense); err != nil {
			t.Errorf("Test Gte reuse for string failed(axb): %v", err)
		}
		if ret, err = a.gteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Gte reuse for string failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}
}

/* Lt */

func TestDense_ltDD_Transitivity(t *testing.T) {
	fI := func(a, b, c *QCDenseI) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}

	fIterI := func(a, b, c *QCDenseI) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b2, AsSameType())
		bxc, _ = b1.ltDD(c2, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.ltDD(b1, AsSameType())
		bxc, _ = b2.ltDD(c1, AsSameType())
		axc, _ = a2.ltDD(c1, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b1, AsSameType())
		bxc, _ = b1.ltDD(c1, AsSameType())
		axc, _ = a1.ltDD(c1, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, b, c *QCDenseI8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}

	fIterI8 := func(a, b, c *QCDenseI8) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b2, AsSameType())
		bxc, _ = b1.ltDD(c2, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.ltDD(b1, AsSameType())
		bxc, _ = b2.ltDD(c1, AsSameType())
		axc, _ = a2.ltDD(c1, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b1, AsSameType())
		bxc, _ = b1.ltDD(c1, AsSameType())
		axc, _ = a1.ltDD(c1, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, b, c *QCDenseI16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}

	fIterI16 := func(a, b, c *QCDenseI16) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b2, AsSameType())
		bxc, _ = b1.ltDD(c2, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.ltDD(b1, AsSameType())
		bxc, _ = b2.ltDD(c1, AsSameType())
		axc, _ = a2.ltDD(c1, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b1, AsSameType())
		bxc, _ = b1.ltDD(c1, AsSameType())
		axc, _ = a1.ltDD(c1, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, b, c *QCDenseI32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}

	fIterI32 := func(a, b, c *QCDenseI32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b2, AsSameType())
		bxc, _ = b1.ltDD(c2, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.ltDD(b1, AsSameType())
		bxc, _ = b2.ltDD(c1, AsSameType())
		axc, _ = a2.ltDD(c1, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b1, AsSameType())
		bxc, _ = b1.ltDD(c1, AsSameType())
		axc, _ = a1.ltDD(c1, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, b, c *QCDenseI64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}

	fIterI64 := func(a, b, c *QCDenseI64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b2, AsSameType())
		bxc, _ = b1.ltDD(c2, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.ltDD(b1, AsSameType())
		bxc, _ = b2.ltDD(c1, AsSameType())
		axc, _ = a2.ltDD(c1, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b1, AsSameType())
		bxc, _ = b1.ltDD(c1, AsSameType())
		axc, _ = a1.ltDD(c1, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, b, c *QCDenseU) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}

	fIterU := func(a, b, c *QCDenseU) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b2, AsSameType())
		bxc, _ = b1.ltDD(c2, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.ltDD(b1, AsSameType())
		bxc, _ = b2.ltDD(c1, AsSameType())
		axc, _ = a2.ltDD(c1, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b1, AsSameType())
		bxc, _ = b1.ltDD(c1, AsSameType())
		axc, _ = a1.ltDD(c1, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, b, c *QCDenseU8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}

	fIterU8 := func(a, b, c *QCDenseU8) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b2, AsSameType())
		bxc, _ = b1.ltDD(c2, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.ltDD(b1, AsSameType())
		bxc, _ = b2.ltDD(c1, AsSameType())
		axc, _ = a2.ltDD(c1, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b1, AsSameType())
		bxc, _ = b1.ltDD(c1, AsSameType())
		axc, _ = a1.ltDD(c1, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, b, c *QCDenseU16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}

	fIterU16 := func(a, b, c *QCDenseU16) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b2, AsSameType())
		bxc, _ = b1.ltDD(c2, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.ltDD(b1, AsSameType())
		bxc, _ = b2.ltDD(c1, AsSameType())
		axc, _ = a2.ltDD(c1, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b1, AsSameType())
		bxc, _ = b1.ltDD(c1, AsSameType())
		axc, _ = a1.ltDD(c1, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, b, c *QCDenseU32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}

	fIterU32 := func(a, b, c *QCDenseU32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b2, AsSameType())
		bxc, _ = b1.ltDD(c2, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.ltDD(b1, AsSameType())
		bxc, _ = b2.ltDD(c1, AsSameType())
		axc, _ = a2.ltDD(c1, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b1, AsSameType())
		bxc, _ = b1.ltDD(c1, AsSameType())
		axc, _ = a1.ltDD(c1, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, b, c *QCDenseU64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}

	fIterU64 := func(a, b, c *QCDenseU64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b2, AsSameType())
		bxc, _ = b1.ltDD(c2, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.ltDD(b1, AsSameType())
		bxc, _ = b2.ltDD(c1, AsSameType())
		axc, _ = a2.ltDD(c1, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b1, AsSameType())
		bxc, _ = b1.ltDD(c1, AsSameType())
		axc, _ = a1.ltDD(c1, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU64, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, b, c *QCDenseF32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for float32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for float32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for float32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}

	fIterF32 := func(a, b, c *QCDenseF32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b2, AsSameType())
		bxc, _ = b1.ltDD(c2, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.ltDD(b1, AsSameType())
		bxc, _ = b2.ltDD(c1, AsSameType())
		axc, _ = a2.ltDD(c1, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b1, AsSameType())
		bxc, _ = b1.ltDD(c1, AsSameType())
		axc, _ = a1.ltDD(c1, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, b, c *QCDenseF64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for float64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for float64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for float64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}

	fIterF64 := func(a, b, c *QCDenseF64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b2, AsSameType())
		bxc, _ = b1.ltDD(c2, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.ltDD(b1, AsSameType())
		bxc, _ = b2.ltDD(c1, AsSameType())
		axc, _ = a2.ltDD(c1, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDD(b1, AsSameType())
		bxc, _ = b1.ltDD(c1, AsSameType())
		axc, _ = a1.ltDD(c1, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF64, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, b, c *QCDenseStr) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt transitivity for string failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for string failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for string failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}

	fIterStr := func(a, b, c *QCDenseStr) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.ltDD(b2)
		bxc, _ = b1.ltDD(c2)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// b iter
		axb, _ = a2.ltDD(b1)
		bxc, _ = b2.ltDD(c1)
		axc, _ = a2.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// both a and b iter
		axb, _ = a1.ltDD(b1)
		bxc, _ = b1.ltDD(c1)
		axc, _ = a1.ltDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterStr, nil); err != nil {
		t.Error(err)
	}
}

func Test_Dense_ltDD_funcOpts(t *testing.T) {
	fI := func(a, b *QCDenseI) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for int failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for int failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lt as same type reuse for int failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt as same type reuse for int failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.ltDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lt for int failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, b *QCDenseI8) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for int8 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for int8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int8, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lt as same type reuse for int8 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt as same type reuse for int8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.ltDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lt for int8 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, b *QCDenseI16) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for int16 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for int16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int16, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lt as same type reuse for int16 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt as same type reuse for int16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.ltDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lt for int16 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, b *QCDenseI32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for int32 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for int32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int32, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lt as same type reuse for int32 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt as same type reuse for int32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.ltDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lt for int32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, b *QCDenseI64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for int64 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for int64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int64, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lt as same type reuse for int64 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt as same type reuse for int64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.ltDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lt for int64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, b *QCDenseU) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for uint failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for uint failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lt as same type reuse for uint failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt as same type reuse for uint failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.ltDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lt for uint failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, b *QCDenseU8) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for uint8 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for uint8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint8, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lt as same type reuse for uint8 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt as same type reuse for uint8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.ltDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lt for uint8 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, b *QCDenseU16) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for uint16 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for uint16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint16, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lt as same type reuse for uint16 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt as same type reuse for uint16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.ltDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lt for uint16 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, b *QCDenseU32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for uint32 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for uint32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint32, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lt as same type reuse for uint32 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt as same type reuse for uint32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.ltDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lt for uint32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, b *QCDenseU64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for uint64 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for uint64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint64, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lt as same type reuse for uint64 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt as same type reuse for uint64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.ltDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lt for uint64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, b *QCDenseF32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for float32 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for float32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Float32, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lt as same type reuse for float32 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt as same type reuse for float32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.ltDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lt for float32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, b *QCDenseF64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for float64 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for float64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Float64, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lt as same type reuse for float64 failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt as same type reuse for float64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.ltDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lt for float64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, b *QCDenseStr) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.ltDD(b.Dense); err != nil {
			t.Errorf("Test Lt reuse for string failed(axb): %v", err)
		}
		if ret, err = a.ltDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lt reuse for string failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}
}

/* Lte */

func TestDense_lteDD_Transitivity(t *testing.T) {
	fI := func(a, b, c *QCDenseI) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}

	fIterI := func(a, b, c *QCDenseI) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b2, AsSameType())
		bxc, _ = b1.lteDD(c2, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.lteDD(b1, AsSameType())
		bxc, _ = b2.lteDD(c1, AsSameType())
		axc, _ = a2.lteDD(c1, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b1, AsSameType())
		bxc, _ = b1.lteDD(c1, AsSameType())
		axc, _ = a1.lteDD(c1, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, b, c *QCDenseI8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}

	fIterI8 := func(a, b, c *QCDenseI8) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b2, AsSameType())
		bxc, _ = b1.lteDD(c2, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.lteDD(b1, AsSameType())
		bxc, _ = b2.lteDD(c1, AsSameType())
		axc, _ = a2.lteDD(c1, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b1, AsSameType())
		bxc, _ = b1.lteDD(c1, AsSameType())
		axc, _ = a1.lteDD(c1, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, b, c *QCDenseI16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}

	fIterI16 := func(a, b, c *QCDenseI16) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b2, AsSameType())
		bxc, _ = b1.lteDD(c2, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.lteDD(b1, AsSameType())
		bxc, _ = b2.lteDD(c1, AsSameType())
		axc, _ = a2.lteDD(c1, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b1, AsSameType())
		bxc, _ = b1.lteDD(c1, AsSameType())
		axc, _ = a1.lteDD(c1, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, b, c *QCDenseI32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}

	fIterI32 := func(a, b, c *QCDenseI32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b2, AsSameType())
		bxc, _ = b1.lteDD(c2, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.lteDD(b1, AsSameType())
		bxc, _ = b2.lteDD(c1, AsSameType())
		axc, _ = a2.lteDD(c1, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b1, AsSameType())
		bxc, _ = b1.lteDD(c1, AsSameType())
		axc, _ = a1.lteDD(c1, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, b, c *QCDenseI64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}

	fIterI64 := func(a, b, c *QCDenseI64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b2, AsSameType())
		bxc, _ = b1.lteDD(c2, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.lteDD(b1, AsSameType())
		bxc, _ = b2.lteDD(c1, AsSameType())
		axc, _ = a2.lteDD(c1, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b1, AsSameType())
		bxc, _ = b1.lteDD(c1, AsSameType())
		axc, _ = a1.lteDD(c1, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, b, c *QCDenseU) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}

	fIterU := func(a, b, c *QCDenseU) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b2, AsSameType())
		bxc, _ = b1.lteDD(c2, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.lteDD(b1, AsSameType())
		bxc, _ = b2.lteDD(c1, AsSameType())
		axc, _ = a2.lteDD(c1, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b1, AsSameType())
		bxc, _ = b1.lteDD(c1, AsSameType())
		axc, _ = a1.lteDD(c1, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, b, c *QCDenseU8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}

	fIterU8 := func(a, b, c *QCDenseU8) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b2, AsSameType())
		bxc, _ = b1.lteDD(c2, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.lteDD(b1, AsSameType())
		bxc, _ = b2.lteDD(c1, AsSameType())
		axc, _ = a2.lteDD(c1, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b1, AsSameType())
		bxc, _ = b1.lteDD(c1, AsSameType())
		axc, _ = a1.lteDD(c1, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, b, c *QCDenseU16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}

	fIterU16 := func(a, b, c *QCDenseU16) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b2, AsSameType())
		bxc, _ = b1.lteDD(c2, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.lteDD(b1, AsSameType())
		bxc, _ = b2.lteDD(c1, AsSameType())
		axc, _ = a2.lteDD(c1, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b1, AsSameType())
		bxc, _ = b1.lteDD(c1, AsSameType())
		axc, _ = a1.lteDD(c1, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, b, c *QCDenseU32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}

	fIterU32 := func(a, b, c *QCDenseU32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b2, AsSameType())
		bxc, _ = b1.lteDD(c2, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.lteDD(b1, AsSameType())
		bxc, _ = b2.lteDD(c1, AsSameType())
		axc, _ = a2.lteDD(c1, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b1, AsSameType())
		bxc, _ = b1.lteDD(c1, AsSameType())
		axc, _ = a1.lteDD(c1, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, b, c *QCDenseU64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}

	fIterU64 := func(a, b, c *QCDenseU64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b2, AsSameType())
		bxc, _ = b1.lteDD(c2, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.lteDD(b1, AsSameType())
		bxc, _ = b2.lteDD(c1, AsSameType())
		axc, _ = a2.lteDD(c1, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b1, AsSameType())
		bxc, _ = b1.lteDD(c1, AsSameType())
		axc, _ = a1.lteDD(c1, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU64, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, b, c *QCDenseF32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for float32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for float32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for float32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}

	fIterF32 := func(a, b, c *QCDenseF32) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b2, AsSameType())
		bxc, _ = b1.lteDD(c2, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.lteDD(b1, AsSameType())
		bxc, _ = b2.lteDD(c1, AsSameType())
		axc, _ = a2.lteDD(c1, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b1, AsSameType())
		bxc, _ = b1.lteDD(c1, AsSameType())
		axc, _ = a1.lteDD(c1, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, b, c *QCDenseF64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for float64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for float64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for float64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}

	fIterF64 := func(a, b, c *QCDenseF64) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b2, AsSameType())
		bxc, _ = b1.lteDD(c2, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// b iter bools
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a2.lteDD(b1, AsSameType())
		bxc, _ = b2.lteDD(c1, AsSameType())
		axc, _ = a2.lteDD(c1, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		// both a and b iter bools
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDD(b1, AsSameType())
		bxc, _ = b1.lteDD(c1, AsSameType())
		axc, _ = a1.lteDD(c1, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF64, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, b, c *QCDenseStr) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte transitivity for string failed (axb) : %v ", err)
			return false
		}
		if bxc, err = b.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for string failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for string failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}

	fIterStr := func(a, b, c *QCDenseStr) bool {
		var axb, bxc, axc *Dense
		var a1, b1, c1 *Dense // sliced
		var a2, b2, c2 *Dense // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		a2 = a1.Materialize().(*Dense)
		b1, _ = sliceDense(a.Dense, makeRS(0, 5))
		b2 = b1.Materialize().(*Dense)
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.lteDD(b2)
		bxc, _ = b1.lteDD(c2)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// b iter
		axb, _ = a2.lteDD(b1)
		bxc, _ = b2.lteDD(c1)
		axc, _ = a2.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// both a and b iter
		axb, _ = a1.lteDD(b1)
		bxc, _ = b1.lteDD(c1)
		axc, _ = a1.lteDD(c1)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterStr, nil); err != nil {
		t.Error(err)
	}
}

func Test_Dense_lteDD_funcOpts(t *testing.T) {
	fI := func(a, b *QCDenseI) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for int failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for int failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lte as same type reuse for int failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte as same type reuse for int failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.lteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lte for int failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, b *QCDenseI8) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for int8 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for int8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int8, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lte as same type reuse for int8 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte as same type reuse for int8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.lteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lte for int8 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, b *QCDenseI16) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for int16 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for int16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int16, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lte as same type reuse for int16 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte as same type reuse for int16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.lteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lte for int16 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, b *QCDenseI32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for int32 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for int32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int32, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lte as same type reuse for int32 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte as same type reuse for int32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.lteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lte for int32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, b *QCDenseI64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for int64 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for int64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Int64, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lte as same type reuse for int64 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte as same type reuse for int64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.lteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lte for int64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, b *QCDenseU) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for uint failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for uint failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lte as same type reuse for uint failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte as same type reuse for uint failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.lteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lte for uint failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, b *QCDenseU8) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for uint8 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for uint8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint8, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lte as same type reuse for uint8 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte as same type reuse for uint8 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.lteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lte for uint8 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, b *QCDenseU16) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for uint16 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for uint16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint16, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lte as same type reuse for uint16 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte as same type reuse for uint16 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.lteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lte for uint16 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, b *QCDenseU32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for uint32 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for uint32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint32, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lte as same type reuse for uint32 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte as same type reuse for uint32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.lteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lte for uint32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, b *QCDenseU64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for uint64 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for uint64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Uint64, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lte as same type reuse for uint64 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte as same type reuse for uint64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.lteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lte for uint64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, b *QCDenseF32) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for float32 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for float32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Float32, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lte as same type reuse for float32 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte as same type reuse for float32 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.lteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lte for float32 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, b *QCDenseF64) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for float64 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for float64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		// reuse as same type
		reuse = recycledDense(Float64, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense, AsSameType()); err != nil {
			t.Errorf("Test Lte as same type reuse for float64 failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, AsSameType(), WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte as same type reuse for float64 failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err == nil {
			t.Error("Expected an error")
			return false
		}

		// unsafe
		if ret, err = a.lteDD(b.Dense, UseUnsafe()); err != nil {
			t.Errorf("Unsafe Lte for float64 failed %v", err)
			return false
		}
		if ret != a.Dense {
			t.Error("Expected ret to be equal to a")
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, b *QCDenseStr) bool {
		var reuse, axb, ret *Dense
		var err error

		reuse = recycledDense(Bool, Shape{a.len()})
		if axb, err = a.lteDD(b.Dense); err != nil {
			t.Errorf("Test Lte reuse for string failed(axb): %v", err)
		}
		if ret, err = a.lteDD(b.Dense, WithReuse(reuse)); err != nil {
			t.Errorf("Test Lte reuse for string failed: %v", err)
			return false
		}
		if ret != reuse {
			t.Errorf("Expected ret to be equal reuse")
			return false
		}
		if !reflect.DeepEqual(axb.Data(), ret.Data()) {
			return false
		}

		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}
}

/* Eq - Dense-Scalar */

func TestDense_eqDS_Transitivity(t *testing.T) {

	fB := func(a, c *QCDenseB, b bool) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for bool failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for bool failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for bool failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fB, nil); err != nil {
		t.Error(err)
	}

	fIterB := func(a, c *QCDenseB, b bool) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterB, nil); err != nil {
		t.Error(err)
	}
	fI := func(a, c *QCDenseI, b int) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for int failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for int failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}

	fIterI := func(a, c *QCDenseI, b int) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, c *QCDenseI8, b int8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for int8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for int8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}

	fIterI8 := func(a, c *QCDenseI8, b int8) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, c *QCDenseI16, b int16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for int16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for int16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}

	fIterI16 := func(a, c *QCDenseI16, b int16) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, c *QCDenseI32, b int32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for int32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for int32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}

	fIterI32 := func(a, c *QCDenseI32, b int32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, c *QCDenseI64, b int64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for int64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for int64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for int64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}

	fIterI64 := func(a, c *QCDenseI64, b int64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, c *QCDenseU, b uint) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for uint failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for uint failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}

	fIterU := func(a, c *QCDenseU, b uint) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, c *QCDenseU8, b uint8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for uint8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for uint8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}

	fIterU8 := func(a, c *QCDenseU8, b uint8) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, c *QCDenseU16, b uint16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for uint16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for uint16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}

	fIterU16 := func(a, c *QCDenseU16, b uint16) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, c *QCDenseU32, b uint32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for uint32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for uint32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}

	fIterU32 := func(a, c *QCDenseU32, b uint32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, c *QCDenseU64, b uint64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for uint64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for uint64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uint64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}

	fIterU64 := func(a, c *QCDenseU64, b uint64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU64, nil); err != nil {
		t.Error(err)
	}
	fUintptr := func(a, c *QCDenseUintptr, b uintptr) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for uintptr failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for uintptr failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for uintptr failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fUintptr, nil); err != nil {
		t.Error(err)
	}

	fIterUintptr := func(a, c *QCDenseUintptr, b uintptr) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterUintptr, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, c *QCDenseF32, b float32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for float32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for float32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for float32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}

	fIterF32 := func(a, c *QCDenseF32, b float32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, c *QCDenseF64, b float64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for float64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for float64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for float64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}

	fIterF64 := func(a, c *QCDenseF64, b float64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF64, nil); err != nil {
		t.Error(err)
	}
	fC64 := func(a, c *QCDenseC64, b complex64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for complex64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for complex64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for complex64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fC64, nil); err != nil {
		t.Error(err)
	}

	fIterC64 := func(a, c *QCDenseC64, b complex64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []complex64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.complex64s()
		bcs = bxc.complex64s()
		acs = axc.complex64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterC64, nil); err != nil {
		t.Error(err)
	}
	fC128 := func(a, c *QCDenseC128, b complex128) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for complex128 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for complex128 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for complex128 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fC128, nil); err != nil {
		t.Error(err)
	}

	fIterC128 := func(a, c *QCDenseC128, b complex128) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []complex128
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.eqDS(b, AsSameType())
		bxc, _ = c2.eqDS(b, AsSameType())
		axc, _ = a1.eqDD(c2, AsSameType())

		abs = axb.complex128s()
		bcs = bxc.complex128s()
		acs = axc.complex128s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterC128, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, c *QCDenseStr, b string) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for string failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.eqDS(b); err != nil {
			t.Errorf("Test Eq transitivity for string failed (bxc): %v", err)
			return false
		}
		if axc, err = a.eqDD(c.Dense); err != nil {
			t.Errorf("Test Eq transitivity for string failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}

	fIterStr := func(a, c *QCDenseStr, b string) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.eqDS(b)
		bxc, _ = c2.eqDS(b)
		axc, _ = a1.eqDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterStr, nil); err != nil {
		t.Error(err)
	}
}

/* Gt - Dense-Scalar */

func TestDense_gtDS_Transitivity(t *testing.T) {

	fI := func(a, c *QCDenseI, b int) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for int failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for int failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}

	fIterI := func(a, c *QCDenseI, b int) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDS(b, AsSameType())
		bxc, _ = c2.ltDS(b, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, c *QCDenseI8, b int8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for int8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for int8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}

	fIterI8 := func(a, c *QCDenseI8, b int8) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDS(b, AsSameType())
		bxc, _ = c2.ltDS(b, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, c *QCDenseI16, b int16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for int16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for int16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}

	fIterI16 := func(a, c *QCDenseI16, b int16) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDS(b, AsSameType())
		bxc, _ = c2.ltDS(b, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, c *QCDenseI32, b int32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for int32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for int32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}

	fIterI32 := func(a, c *QCDenseI32, b int32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDS(b, AsSameType())
		bxc, _ = c2.ltDS(b, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, c *QCDenseI64, b int64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for int64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for int64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for int64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}

	fIterI64 := func(a, c *QCDenseI64, b int64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDS(b, AsSameType())
		bxc, _ = c2.ltDS(b, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, c *QCDenseU, b uint) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for uint failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for uint failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}

	fIterU := func(a, c *QCDenseU, b uint) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDS(b, AsSameType())
		bxc, _ = c2.ltDS(b, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, c *QCDenseU8, b uint8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for uint8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for uint8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}

	fIterU8 := func(a, c *QCDenseU8, b uint8) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDS(b, AsSameType())
		bxc, _ = c2.ltDS(b, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, c *QCDenseU16, b uint16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for uint16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for uint16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}

	fIterU16 := func(a, c *QCDenseU16, b uint16) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDS(b, AsSameType())
		bxc, _ = c2.ltDS(b, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, c *QCDenseU32, b uint32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for uint32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for uint32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}

	fIterU32 := func(a, c *QCDenseU32, b uint32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDS(b, AsSameType())
		bxc, _ = c2.ltDS(b, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, c *QCDenseU64, b uint64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for uint64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for uint64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for uint64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}

	fIterU64 := func(a, c *QCDenseU64, b uint64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDS(b, AsSameType())
		bxc, _ = c2.ltDS(b, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU64, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, c *QCDenseF32, b float32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for float32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for float32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for float32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}

	fIterF32 := func(a, c *QCDenseF32, b float32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDS(b, AsSameType())
		bxc, _ = c2.ltDS(b, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, c *QCDenseF64, b float64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for float64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for float64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for float64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}

	fIterF64 := func(a, c *QCDenseF64, b float64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gtDS(b, AsSameType())
		bxc, _ = c2.ltDS(b, AsSameType())
		axc, _ = a1.gtDD(c2, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF64, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, c *QCDenseStr, b string) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gtDS(b); err != nil {
			t.Errorf("Test Gt transitivity for string failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.ltDS(b); err != nil {
			t.Errorf("Test Gt transitivity for string failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gtDD(c.Dense); err != nil {
			t.Errorf("Test Gt transitivity for string failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}

	fIterStr := func(a, c *QCDenseStr, b string) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.gtDS(b)
		bxc, _ = c2.ltDS(b)
		axc, _ = a1.gtDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterStr, nil); err != nil {
		t.Error(err)
	}
}

/* Gte - Dense-Scalar */

func TestDense_gteDS_Transitivity(t *testing.T) {

	fI := func(a, c *QCDenseI, b int) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for int failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for int failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}

	fIterI := func(a, c *QCDenseI, b int) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDS(b, AsSameType())
		bxc, _ = c2.lteDS(b, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, c *QCDenseI8, b int8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for int8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for int8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}

	fIterI8 := func(a, c *QCDenseI8, b int8) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDS(b, AsSameType())
		bxc, _ = c2.lteDS(b, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, c *QCDenseI16, b int16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for int16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for int16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}

	fIterI16 := func(a, c *QCDenseI16, b int16) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDS(b, AsSameType())
		bxc, _ = c2.lteDS(b, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, c *QCDenseI32, b int32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for int32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for int32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}

	fIterI32 := func(a, c *QCDenseI32, b int32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDS(b, AsSameType())
		bxc, _ = c2.lteDS(b, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, c *QCDenseI64, b int64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for int64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for int64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for int64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}

	fIterI64 := func(a, c *QCDenseI64, b int64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDS(b, AsSameType())
		bxc, _ = c2.lteDS(b, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, c *QCDenseU, b uint) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for uint failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for uint failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}

	fIterU := func(a, c *QCDenseU, b uint) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDS(b, AsSameType())
		bxc, _ = c2.lteDS(b, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, c *QCDenseU8, b uint8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for uint8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for uint8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}

	fIterU8 := func(a, c *QCDenseU8, b uint8) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDS(b, AsSameType())
		bxc, _ = c2.lteDS(b, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, c *QCDenseU16, b uint16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for uint16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for uint16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}

	fIterU16 := func(a, c *QCDenseU16, b uint16) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDS(b, AsSameType())
		bxc, _ = c2.lteDS(b, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, c *QCDenseU32, b uint32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for uint32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for uint32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}

	fIterU32 := func(a, c *QCDenseU32, b uint32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDS(b, AsSameType())
		bxc, _ = c2.lteDS(b, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, c *QCDenseU64, b uint64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for uint64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for uint64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for uint64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}

	fIterU64 := func(a, c *QCDenseU64, b uint64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDS(b, AsSameType())
		bxc, _ = c2.lteDS(b, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU64, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, c *QCDenseF32, b float32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for float32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for float32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for float32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}

	fIterF32 := func(a, c *QCDenseF32, b float32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDS(b, AsSameType())
		bxc, _ = c2.lteDS(b, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, c *QCDenseF64, b float64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for float64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for float64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for float64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}

	fIterF64 := func(a, c *QCDenseF64, b float64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.gteDS(b, AsSameType())
		bxc, _ = c2.lteDS(b, AsSameType())
		axc, _ = a1.gteDD(c2, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF64, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, c *QCDenseStr, b string) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.gteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for string failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.lteDS(b); err != nil {
			t.Errorf("Test Gte transitivity for string failed (bxc): %v", err)
			return false
		}
		if axc, err = a.gteDD(c.Dense); err != nil {
			t.Errorf("Test Gte transitivity for string failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}

	fIterStr := func(a, c *QCDenseStr, b string) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.gteDS(b)
		bxc, _ = c2.lteDS(b)
		axc, _ = a1.gteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterStr, nil); err != nil {
		t.Error(err)
	}
}

/* Lt - Dense-Scalar */

func TestDense_ltDS_Transitivity(t *testing.T) {

	fI := func(a, c *QCDenseI, b int) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for int failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for int failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}

	fIterI := func(a, c *QCDenseI, b int) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDS(b, AsSameType())
		bxc, _ = c2.gtDS(b, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, c *QCDenseI8, b int8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for int8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for int8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}

	fIterI8 := func(a, c *QCDenseI8, b int8) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDS(b, AsSameType())
		bxc, _ = c2.gtDS(b, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, c *QCDenseI16, b int16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for int16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for int16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}

	fIterI16 := func(a, c *QCDenseI16, b int16) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDS(b, AsSameType())
		bxc, _ = c2.gtDS(b, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, c *QCDenseI32, b int32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for int32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for int32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}

	fIterI32 := func(a, c *QCDenseI32, b int32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDS(b, AsSameType())
		bxc, _ = c2.gtDS(b, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, c *QCDenseI64, b int64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for int64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for int64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for int64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}

	fIterI64 := func(a, c *QCDenseI64, b int64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDS(b, AsSameType())
		bxc, _ = c2.gtDS(b, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, c *QCDenseU, b uint) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for uint failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for uint failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}

	fIterU := func(a, c *QCDenseU, b uint) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDS(b, AsSameType())
		bxc, _ = c2.gtDS(b, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, c *QCDenseU8, b uint8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for uint8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for uint8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}

	fIterU8 := func(a, c *QCDenseU8, b uint8) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDS(b, AsSameType())
		bxc, _ = c2.gtDS(b, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, c *QCDenseU16, b uint16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for uint16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for uint16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}

	fIterU16 := func(a, c *QCDenseU16, b uint16) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDS(b, AsSameType())
		bxc, _ = c2.gtDS(b, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, c *QCDenseU32, b uint32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for uint32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for uint32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}

	fIterU32 := func(a, c *QCDenseU32, b uint32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDS(b, AsSameType())
		bxc, _ = c2.gtDS(b, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, c *QCDenseU64, b uint64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for uint64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for uint64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for uint64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}

	fIterU64 := func(a, c *QCDenseU64, b uint64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDS(b, AsSameType())
		bxc, _ = c2.gtDS(b, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU64, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, c *QCDenseF32, b float32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for float32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for float32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for float32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}

	fIterF32 := func(a, c *QCDenseF32, b float32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDS(b, AsSameType())
		bxc, _ = c2.gtDS(b, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, c *QCDenseF64, b float64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for float64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for float64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for float64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}

	fIterF64 := func(a, c *QCDenseF64, b float64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.ltDS(b, AsSameType())
		bxc, _ = c2.gtDS(b, AsSameType())
		axc, _ = a1.ltDD(c2, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF64, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, c *QCDenseStr, b string) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.ltDS(b); err != nil {
			t.Errorf("Test Lt transitivity for string failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gtDS(b); err != nil {
			t.Errorf("Test Lt transitivity for string failed (bxc): %v", err)
			return false
		}
		if axc, err = a.ltDD(c.Dense); err != nil {
			t.Errorf("Test Lt transitivity for string failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}

	fIterStr := func(a, c *QCDenseStr, b string) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.ltDS(b)
		bxc, _ = c2.gtDS(b)
		axc, _ = a1.ltDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterStr, nil); err != nil {
		t.Error(err)
	}
}

/* Lte - Dense-Scalar */

func TestDense_lteDS_Transitivity(t *testing.T) {

	fI := func(a, c *QCDenseI, b int) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for int failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for int failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI, nil); err != nil {
		t.Error(err)
	}

	fIterI := func(a, c *QCDenseI, b int) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDS(b, AsSameType())
		bxc, _ = c2.gteDS(b, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.ints()
		bcs = bxc.ints()
		acs = axc.ints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI, nil); err != nil {
		t.Error(err)
	}
	fI8 := func(a, c *QCDenseI8, b int8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for int8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for int8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI8, nil); err != nil {
		t.Error(err)
	}

	fIterI8 := func(a, c *QCDenseI8, b int8) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDS(b, AsSameType())
		bxc, _ = c2.gteDS(b, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.int8s()
		bcs = bxc.int8s()
		acs = axc.int8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI8, nil); err != nil {
		t.Error(err)
	}
	fI16 := func(a, c *QCDenseI16, b int16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for int16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for int16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI16, nil); err != nil {
		t.Error(err)
	}

	fIterI16 := func(a, c *QCDenseI16, b int16) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDS(b, AsSameType())
		bxc, _ = c2.gteDS(b, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.int16s()
		bcs = bxc.int16s()
		acs = axc.int16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI16, nil); err != nil {
		t.Error(err)
	}
	fI32 := func(a, c *QCDenseI32, b int32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for int32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for int32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI32, nil); err != nil {
		t.Error(err)
	}

	fIterI32 := func(a, c *QCDenseI32, b int32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDS(b, AsSameType())
		bxc, _ = c2.gteDS(b, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.int32s()
		bcs = bxc.int32s()
		acs = axc.int32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI32, nil); err != nil {
		t.Error(err)
	}
	fI64 := func(a, c *QCDenseI64, b int64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for int64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for int64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for int64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fI64, nil); err != nil {
		t.Error(err)
	}

	fIterI64 := func(a, c *QCDenseI64, b int64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []int64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDS(b, AsSameType())
		bxc, _ = c2.gteDS(b, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.int64s()
		bcs = bxc.int64s()
		acs = axc.int64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterI64, nil); err != nil {
		t.Error(err)
	}
	fU := func(a, c *QCDenseU, b uint) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for uint failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for uint failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU, nil); err != nil {
		t.Error(err)
	}

	fIterU := func(a, c *QCDenseU, b uint) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDS(b, AsSameType())
		bxc, _ = c2.gteDS(b, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.uints()
		bcs = bxc.uints()
		acs = axc.uints()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU, nil); err != nil {
		t.Error(err)
	}
	fU8 := func(a, c *QCDenseU8, b uint8) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for uint8 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for uint8 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint8 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU8, nil); err != nil {
		t.Error(err)
	}

	fIterU8 := func(a, c *QCDenseU8, b uint8) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint8
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDS(b, AsSameType())
		bxc, _ = c2.gteDS(b, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.uint8s()
		bcs = bxc.uint8s()
		acs = axc.uint8s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU8, nil); err != nil {
		t.Error(err)
	}
	fU16 := func(a, c *QCDenseU16, b uint16) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for uint16 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for uint16 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint16 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU16, nil); err != nil {
		t.Error(err)
	}

	fIterU16 := func(a, c *QCDenseU16, b uint16) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint16
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDS(b, AsSameType())
		bxc, _ = c2.gteDS(b, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.uint16s()
		bcs = bxc.uint16s()
		acs = axc.uint16s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU16, nil); err != nil {
		t.Error(err)
	}
	fU32 := func(a, c *QCDenseU32, b uint32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for uint32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for uint32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU32, nil); err != nil {
		t.Error(err)
	}

	fIterU32 := func(a, c *QCDenseU32, b uint32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDS(b, AsSameType())
		bxc, _ = c2.gteDS(b, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.uint32s()
		bcs = bxc.uint32s()
		acs = axc.uint32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU32, nil); err != nil {
		t.Error(err)
	}
	fU64 := func(a, c *QCDenseU64, b uint64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for uint64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for uint64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for uint64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fU64, nil); err != nil {
		t.Error(err)
	}

	fIterU64 := func(a, c *QCDenseU64, b uint64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []uint64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDS(b, AsSameType())
		bxc, _ = c2.gteDS(b, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.uint64s()
		bcs = bxc.uint64s()
		acs = axc.uint64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterU64, nil); err != nil {
		t.Error(err)
	}
	fF32 := func(a, c *QCDenseF32, b float32) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for float32 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for float32 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for float32 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF32, nil); err != nil {
		t.Error(err)
	}

	fIterF32 := func(a, c *QCDenseF32, b float32) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float32
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDS(b, AsSameType())
		bxc, _ = c2.gteDS(b, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.float32s()
		bcs = bxc.float32s()
		acs = axc.float32s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF32, nil); err != nil {
		t.Error(err)
	}
	fF64 := func(a, c *QCDenseF64, b float64) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for float64 failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for float64 failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for float64 failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fF64, nil); err != nil {
		t.Error(err)
	}

	fIterF64 := func(a, c *QCDenseF64, b float64) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		var abs, bcs, acs []float64
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter bools
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		// a iter asSame
		axb, _ = a1.lteDS(b, AsSameType())
		bxc, _ = c2.gteDS(b, AsSameType())
		axc, _ = a1.lteDD(c2, AsSameType())

		abs = axb.float64s()
		bcs = bxc.float64s()
		acs = axc.float64s()

		for i, vab := range abs {
			if vab == 1 && bcs[i] == 1 {
				if acs[i] != 1 {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fIterF64, nil); err != nil {
		t.Error(err)
	}
	fStr := func(a, c *QCDenseStr, b string) bool {
		var axb, bxc, axc *Dense
		var err error
		if axb, err = a.lteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for string failed (axb) : %v ", err)
			return false
		}
		if bxc, err = c.gteDS(b); err != nil {
			t.Errorf("Test Lte transitivity for string failed (bxc): %v", err)
			return false
		}
		if axc, err = a.lteDD(c.Dense); err != nil {
			t.Errorf("Test Lte transitivity for string failed (axc): %v", err)
			return false
		}

		ab := axb.bools()
		bc := bxc.bools()
		ac := axc.bools()

		for i, vab := range ab {
			if vab && bc[i] {
				if !ac[i] {
					return false
				}
			}
		}
		return true
	}
	if err := quick.Check(fStr, nil); err != nil {
		t.Error(err)
	}

	fIterStr := func(a, c *QCDenseStr, b string) bool {
		var axb, bxc, axc *Dense
		var a1, c1 *Dense // sliced
		var c2 *Dense     // materialized slice

		var abb, bcb, acb []bool
		// set up
		a1, _ = sliceDense(a.Dense, makeRS(0, 5))
		c1, _ = sliceDense(c.Dense, makeRS(0, 5))
		c2 = c1.Materialize().(*Dense)

		// a iter
		axb, _ = a1.lteDS(b)
		bxc, _ = c2.gteDS(b)
		axc, _ = a1.lteDD(c2)

		abb = axb.bools()
		bcb = bxc.bools()
		acb = axc.bools()

		for i, vab := range abb {
			if vab && bcb[i] {
				if !acb[i] {
					return false
				}
			}
		}

		return true
	}
	if err := quick.Check(fIterStr, nil); err != nil {
		t.Error(err)
	}
}
