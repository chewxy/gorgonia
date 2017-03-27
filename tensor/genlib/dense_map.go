package main

import (
	"io"
	"text/template"
)

const mapperRaw = `func (t *Dense) mapFn(fn interface{}, incr bool) (err error) {
	switch t.t.Kind() {
	{{range .Kinds -}}
	{{if isParameterized . -}}
	{{else -}}
	case reflect.{{reflectKind .}}:
		if f, ok := fn.(func({{asType .}}){{asType .}}); ok {
			data := t.{{sliceOf .}}
			for i, v := range data {
				{{if isNumber . -}}
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				{{else -}}
					data[i] = f(v)
				{{end -}}
			}
			return nil
		}
		return errors.Errorf(extractionFail, "func({{asType .}}) {{asType .}}", fn)
	{{end -}}
	{{end -}}
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

func (t *Dense) iterMap(fn interface{}, it *FlatIterator, incr bool) (err error) {
	switch t.t.Kind() {
	{{range .Kinds -}}
	{{if isParameterized . -}}
	{{else -}}
	case reflect.{{reflectKind .}}:
		if f, ok := fn.(func({{asType .}}){{asType .}}); ok {
			data := t.{{sliceOf .}}
			var i int
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				v := data[i]
				{{if isNumber . -}}
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				{{else -}}
					data[i] = f(v)
				{{end -}}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func({{asType .}}) {{asType .}}", fn)
	{{end -}}
	{{end -}}
	default:
		// TODO: fix to handle incr
		var f reflect.Value
		var fnT reflect.Type
		if f, fnT, err = reductionFnType(fn, t.t.Type); err != nil {
			return 
		}
		args := make([]reflect.Value, 0, fnT.NumIn())
		var i int
		for i, err = it.Next(); err == nil; i, err = it.Next() {
			args = append(args, reflect.ValueOf(t.Get(i)))
			t.Set(i, f.Call(args)[0].Interface())
			args = args[:0]
		}
		return handleNoOp(err)
	}
	return nil
}

func (t *Dense) iterMapMasked(fn interface{}, it *MultIterator, incr bool) (err error) {
	if !t.IsMasked(){
		return errors.Errorf(maskRequired, "iterMapMasked")
	}
	switch t.t.Kind() {
	{{range .Kinds -}}
	{{if isParameterized . -}}
	{{else -}}
	case reflect.{{reflectKind .}}:
		if f, ok := fn.(func({{asType .}}){{asType .}}); ok {
			data := t.{{sliceOf .}}
			var i int
			for i, err = it.Next(); err == nil; i, err = it.Next() {
				j := it.LastMaskIndex(0)
				v := data[i]
				if !t.mask[j] {
				{{if isNumber . -}}
					if incr {
						data[i] += f(v)
					} else {
						data[i] = f(v)
					}
				{{else -}}
					data[i] = f(v)
				{{end -}}
				}
			}
			return handleNoOp(err)
		}
		return errors.Errorf(extractionFail, "func({{asType .}}) {{asType .}}", fn)
	{{end -}}
	{{end -}}
	default:
		// TODO: fix to handle incr
		var f reflect.Value
		var fnT reflect.Type
		if f, fnT, err = reductionFnType(fn, t.t.Type); err != nil {
			return 
		}
		args := make([]reflect.Value, 0, fnT.NumIn())
		var i int
		for i, err = it.Next(); err == nil; i, err = it.Next() {
			j := it.LastMaskIndex(0)
			if !t.mask[j] {
				args = append(args, reflect.ValueOf(t.Get(i)))
				t.Set(i, f.Call(args)[0].Interface())
				args = args[:0]
			}
		}
		return handleNoOp(err)
	}
	return nil
}
`

var (
	mapper *template.Template
)

func init() {
	mapper = template.Must(template.New("mapper").Funcs(funcs).Parse(mapperRaw))
}

func generateDenseMapper(f io.Writer, generic *ManyKinds) {
	mapper.Execute(f, generic)
}
