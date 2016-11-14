package gorgonia

import "github.com/chewxy/hm"

func isScalarType(t hm.Type) bool {
	switch tt := t.(type) {
	case Dtype:
		return true
	case hm.TypeVariable:
		return isScalarType(hm.Prune(tt))
	default:
		return false
	}
}

func dtypeOf(t hm.Type) (retVal Dtype, err error) {
	pruned := hm.Prune(t)
	switch p := pruned.(type) {
	case Dtype:
		retVal = p
	case *TensorType:
		return dtypeOf(p.of)
	case hm.TypeVariable:
		if p.Instance() == nil {
			err = NewError(typeError, "instance %v does not have a dtype", p)
		}

		return dtypeOf(p.Instance())
	default:
		err = NewError(NotYetImplemented, "dtypeOf of %v not yet implemented", t)
		return
	}

	return

}

func runtimeTypeCheck(expected, got hm.Types) (of Dtype, err error) {
	if len(expected) != len(got) {
		err = NewError(RuntimeError, "Input length mismatch")
		return
	}

	if of, err = dtypeOf(expected[0]); err != nil {
		return
	}

	for i, e := range expected {
		g := got[i]
		if !e.Eq(g) {
			err = NewError(RuntimeError, "Expected input[%d] to be %v. Got %v instead", i, e, got[i])
			return
		}

		if i > 0 {
			var gdt Dtype
			if gdt, err = dtypeOf(g); err == nil {
				if gdt != of {
					err = NewError(RuntimeError, "Different dtypes encountered... Expected %v. Got %v instead", of, gdt)
					return
				}
			} else {
				return
			}
		}
	}
	return
}
