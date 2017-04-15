package tensor

import (
	"reflect"
	"unsafe"

	"github.com/pkg/errors"
)

/*
GENERATED FILE. DO NOT EDIT
*/

func (t *Dense) mapFn(fn interface{}, incr bool) (err error) {
	switch t.t.Kind() {
	case reflect.Bool:
		if f, ok := fn.(func(bool) bool); ok {
			data := t.bools()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						data[i] = f(v)
					}
				}
			} else {
				for i, v := range data {
					data[i] = f(v)
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(bool) bool", fn)
	case reflect.Int:
		if f, ok := fn.(func(int) int); ok {
			data := t.ints()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(int) int", fn)
	case reflect.Int8:
		if f, ok := fn.(func(int8) int8); ok {
			data := t.int8s()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(int8) int8", fn)
	case reflect.Int16:
		if f, ok := fn.(func(int16) int16); ok {
			data := t.int16s()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(int16) int16", fn)
	case reflect.Int32:
		if f, ok := fn.(func(int32) int32); ok {
			data := t.int32s()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(int32) int32", fn)
	case reflect.Int64:
		if f, ok := fn.(func(int64) int64); ok {
			data := t.int64s()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(int64) int64", fn)
	case reflect.Uint:
		if f, ok := fn.(func(uint) uint); ok {
			data := t.uints()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(uint) uint", fn)
	case reflect.Uint8:
		if f, ok := fn.(func(uint8) uint8); ok {
			data := t.uint8s()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(uint8) uint8", fn)
	case reflect.Uint16:
		if f, ok := fn.(func(uint16) uint16); ok {
			data := t.uint16s()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(uint16) uint16", fn)
	case reflect.Uint32:
		if f, ok := fn.(func(uint32) uint32); ok {
			data := t.uint32s()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(uint32) uint32", fn)
	case reflect.Uint64:
		if f, ok := fn.(func(uint64) uint64); ok {
			data := t.uint64s()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(uint64) uint64", fn)
	case reflect.Uintptr:
		if f, ok := fn.(func(uintptr) uintptr); ok {
			data := t.uintptrs()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						data[i] = f(v)
					}
				}
			} else {
				for i, v := range data {
					data[i] = f(v)
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(uintptr) uintptr", fn)
	case reflect.Float32:
		if f, ok := fn.(func(float32) float32); ok {
			data := t.float32s()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(float32) float32", fn)
	case reflect.Float64:
		if f, ok := fn.(func(float64) float64); ok {
			data := t.float64s()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(float64) float64", fn)
	case reflect.Complex64:
		if f, ok := fn.(func(complex64) complex64); ok {
			data := t.complex64s()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(complex64) complex64", fn)
	case reflect.Complex128:
		if f, ok := fn.(func(complex128) complex128); ok {
			data := t.complex128s()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						if !mask[i] {
							if incr {
								data[i] += f(v)
							} else {
								data[i] = f(v)
							}
						}
					}
				}
			} else {
				for i, v := range data {
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(complex128) complex128", fn)
	case reflect.String:
		if f, ok := fn.(func(string) string); ok {
			data := t.strings()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						data[i] = f(v)
					}
				}
			} else {
				for i, v := range data {
					data[i] = f(v)
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(string) string", fn)
	case reflect.UnsafePointer:
		if f, ok := fn.(func(unsafe.Pointer) unsafe.Pointer); ok {
			data := t.unsafePointers()
			if t.IsMasked() {
				mask := t.mask
				if len(mask) == len(data) {
					for i, v := range data {
						data[i] = f(v)
					}
				}
			} else {
				for i, v := range data {
					data[i] = f(v)
				}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func(unsafe.Pointer) unsafe.Pointer", fn)
	default:
		// TODO: fix to handle incr
		var f reflect.Value
		var fnT reflect.Type
		if f, fnT, err = reductionFnType(fn, t.t.Type); err != nil {
			return
		}
		args := make([]reflect.Value, 0, fnT.NumIn())
		for i := 0; i < t.len(); i++ {
			args = append(args, reflect.ValueOf(t.Get(i)))
			t.Set(i, f.Call(args)[0].Interface())
			args = args[:0]
		}
	}
	return nil
}

func (t *Dense) iterMap(fn interface{}, it Iterator, incr bool) (err error) {
	switch t.t.Kind() {
	case reflect.Bool:
		if f, ok := fn.(func(bool) bool); ok {
			data := t.bools()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				data[i] = f(v)
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(bool) bool", fn)
	case reflect.Int:
		if f, ok := fn.(func(int) int); ok {
			data := t.ints()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(int) int", fn)
	case reflect.Int8:
		if f, ok := fn.(func(int8) int8); ok {
			data := t.int8s()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(int8) int8", fn)
	case reflect.Int16:
		if f, ok := fn.(func(int16) int16); ok {
			data := t.int16s()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(int16) int16", fn)
	case reflect.Int32:
		if f, ok := fn.(func(int32) int32); ok {
			data := t.int32s()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(int32) int32", fn)
	case reflect.Int64:
		if f, ok := fn.(func(int64) int64); ok {
			data := t.int64s()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(int64) int64", fn)
	case reflect.Uint:
		if f, ok := fn.(func(uint) uint); ok {
			data := t.uints()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(uint) uint", fn)
	case reflect.Uint8:
		if f, ok := fn.(func(uint8) uint8); ok {
			data := t.uint8s()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(uint8) uint8", fn)
	case reflect.Uint16:
		if f, ok := fn.(func(uint16) uint16); ok {
			data := t.uint16s()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(uint16) uint16", fn)
	case reflect.Uint32:
		if f, ok := fn.(func(uint32) uint32); ok {
			data := t.uint32s()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(uint32) uint32", fn)
	case reflect.Uint64:
		if f, ok := fn.(func(uint64) uint64); ok {
			data := t.uint64s()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(uint64) uint64", fn)
	case reflect.Uintptr:
		if f, ok := fn.(func(uintptr) uintptr); ok {
			data := t.uintptrs()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				data[i] = f(v)
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(uintptr) uintptr", fn)
	case reflect.Float32:
		if f, ok := fn.(func(float32) float32); ok {
			data := t.float32s()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(float32) float32", fn)
	case reflect.Float64:
		if f, ok := fn.(func(float64) float64); ok {
			data := t.float64s()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(float64) float64", fn)
	case reflect.Complex64:
		if f, ok := fn.(func(complex64) complex64); ok {
			data := t.complex64s()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(complex64) complex64", fn)
	case reflect.Complex128:
		if f, ok := fn.(func(complex128) complex128); ok {
			data := t.complex128s()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				if incr {
					data[i] += f(v)
				} else {
					data[i] = f(v)
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(complex128) complex128", fn)
	case reflect.String:
		if f, ok := fn.(func(string) string); ok {
			data := t.strings()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				data[i] = f(v)
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(string) string", fn)
	case reflect.UnsafePointer:
		if f, ok := fn.(func(unsafe.Pointer) unsafe.Pointer); ok {
			data := t.unsafePointers()
			var i int
			for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
				v := data[i]
				data[i] = f(v)
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func(unsafe.Pointer) unsafe.Pointer", fn)
	default:
		// TODO: fix to handle incr
		var f reflect.Value
		var fnT reflect.Type
		if f, fnT, err = reductionFnType(fn, t.t.Type); err != nil {
			return
		}
		args := make([]reflect.Value, 0, fnT.NumIn())
		var i int
		for i, _, err = it.NextValid(); err == nil; i, _, err = it.NextValid() {
			args = append(args, reflect.ValueOf(t.Get(i)))
			t.Set(i, f.Call(args)[0].Interface())
			args = args[:0]
		}
		return handleNoOp(err)
	}
	return nil
}
