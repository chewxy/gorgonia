package main

import (
	"fmt"
	"io"
	"log"
	"text/template"
)

type BinOp struct {
	ArrayType
	OpName string
	OpSymb string
	IsFunc bool
}

var binOps = []struct {
	OpName string
	OpSymb string

	IsFunc bool
}{
	{"Add", "+", false},
	{"Sub", "-", false},
	{"Mul", "*", false},
	{"Div", "/", false},
	{"Pow", "math.Pow", true},
}

var vecscalarOps = []struct {
	OpName string
	OpSymb string

	IsFunc bool
}{
	{"Trans", "+", false},
	{"TransInv", "-", false},
	{"TransInvR", "-", false},
	{"Scale", "*", false},
	{"ScaleInv", "/", false},
	{"ScaleInvR", "/", false},
	{"PowOf", "math.Pow", true},
	{"PowOfR", "math.Pow", true},
}

const binOpRaw = `func (a {{.Name}}) {{.OpName}}(other Number) error {
	b, err := get{{title .Of}}s(other)
	if err != nil {
		return errors.Wrapf(err, opFail, "{{.OpName}}")
	}
	
	if len(a) != len(b){
		return errors.Errorf(lenMismatch,  len(a), len(b))
	}

	{{if ne .VecPkg "" -}}
		{{.VecPkg}}.{{.OpName}}([]{{.Of}}(a), b)
	{{else -}}
		{{$scaleInv := hasPrefix .OpName "ScaleInv" -}}
		{{$div := hasPrefix .OpName "Div" -}}
		{{if or $scaleInv $div -}}var errs errorIndices{{end}}
		for i, v := range b {
			{{if or $scaleInv $div -}}
			if v == {{.Of}}(0) {
				errs = append(errs, i)
				a[i] = 0
				continue
			}

			{{end -}}
			{{if .IsFunc -}}
				a[i] = {{.Of}}({{.OpSymb}} (float64(a[i]), float64(v)))
			{{else -}}
				a[i] {{.OpSymb}}= v 
			{{end -}}
		}

		{{if or $scaleInv $div -}}
			if errs != nil {
				return errs
			}
		{{end -}}
	{{end -}}
	return nil
}
`

const vecScalarOpRaw = `func (a {{.Name}}) {{.OpName}}(other interface{}) (err error) {
	var b {{.Of}}
	if b, err = get{{title .Of}}(other); err != nil{
		return errors.Wrapf(err, opFail, "{{.OpName}}")
	}

	{{if ne .VecPkg "" -}}
	{{.VecPkg}}.{{.OpName}}([]{{.Of}}(a), b)
	{{else -}}
	{{$scaleInv := hasPrefix .OpName "ScaleInv" -}}
	{{$div := hasPrefix .OpName "Div" -}}
	{{if or $scaleInv $div -}}var errs errorIndices{{end}}
		for i, v := range a {
			{{if or $scaleInv $div -}}
			if v == {{.Of}}(0) {
				errs = append(errs, i)
				a[i] = 0 
				continue
			}
			{{end -}}

			a[i] = {{if hasSuffix .OpName "R" -}}
				{{if .IsFunc -}} 
					{{.Of}}({{.OpSymb}}(float64(b), float64(v)))
				{{else -}} 
					b {{.OpSymb}} v 
				{{end -}}
			{{else -}} 
				{{if .IsFunc -}} 
					{{.Of}}({{.OpSymb}}(float64(v), float64(b)))
				{{else -}} 
					v {{.OpSymb}} b
				{{end -}}
			{{end -}}
		}
		{{if or $scaleInv $div -}}
			if errs != nil {
				return errs
			}
		{{end -}}
	{{end -}}
	return nil
}
`
const binOpTestRaw = `func Test_{{.Name}}_{{.OpName}}(t *testing.T){
	a, b, c, _ := prep{{.Name}}Test()

	correct := make({{.Name}}, len(a))
	for i, v := range a {
		correct[i] = {{if .IsFunc -}}
			{{.Of}}({{.OpSymb}}(float64(v), float64(b[i]))) 
		{{else -}} 
			v {{.OpSymb}} b[i] 
		{{end -}}
	}

	// same type
	if err := a.{{.OpName}}(b); err != nil {
		t.Error(err)
	}

	for i, v := range a {
		if v != correct[i]{
			t.Errorf("{{.OpName}} is incorrect. Expected %v. Got %v", correct[i], v)
			break
		}
	}

	// compatible type
	a = {{.Name}}{ {{.TestData0}} }
	if err := a.{{.OpName}}(c); err != nil{
		t.Error(err)
	}

	for i, v := range a {
		if v != correct[i]{
			t.Errorf("{{.OpName}} is incorrect. Expected %v. Got %v", correct[i], v)
			break
		}
	}
	{{$div := hasPrefix .OpName "Div"}}
	{{$scaleInv := hasPrefix .OpName "ScaleInv"}}
	{{if or $div $scaleInv -}}
		{{if hasPrefix .Of "float" -}}
		{{else -}}
			// additional tests for ScaleInv just for completeness sake
			b = {{.Name}}{ {{.TestData0}} }
			if err := a.{{.OpName}}(b); err == nil {
				t.Error("Expected an errrorIndices")
			}

		{{end -}}
	{{end -}}

	// idiotsville 1
	if err := a.{{.OpName}}(b[:3]); err == nil {
		t.Error("Expected an error when performing {{.OpName}} on differing lengths")
	}

	// idiotsville 2
	{{if eq .Name "f64s"}}if err := a.{{.OpName}}(f32s{}); err == nil {{else}}if err := a.{{.OpName}}(f64s{}); err == nil {{end -}} {
		t.Errorf("Expected an error when performing {{.OpName}} on a non-compatible type")
	}
}
`

const vecScalarTestRaw = `func Test_{{.Name}}_{{.OpName}}(t *testing.T){
	a, _, _, b := prep{{.Name}}Test()

	correct := make({{.Name}}, len(a))
	for i, v := range a {
		{{if hasPrefix .OpName "ScaleInv" -}}
			if v == {{.Of}}(0) {
				correct[i] = 0
				continue
			}
		{{end -}}

		correct[i] = {{if hasSuffix .OpName "R" -}}
			{{if .IsFunc -}} {{.Of}}({{.OpSymb}}(float64(b), float64(v)))
			{{else -}} b {{.OpSymb}} v 
			{{end -}}
		{{else -}} 
			{{if .IsFunc -}} {{.Of}}({{.OpSymb}}(float64(v), float64(b)))
			{{else -}} v {{.OpSymb}} b
			{{end -}}
		{{end -}}
	}

	{{if hasPrefix .OpName "ScaleInv" -}}
		{{if hasPrefix .Of "float" -}}
			if err := a.{{.OpName}}(b); err != nil{
				t.Fatal(err)
			}
		{{else -}}
			err := a.{{.OpName}}(b)
			if err == nil {
				t.Error("Expected error (division by zero)")
			}
			if _, ok := err.(errorIndices); !ok{
				t.Fatal(err)
			}
		{{end -}}
	{{else -}}
	if err := a.{{.OpName}}(b); err != nil {
		t.Fatal(err)
	}
	{{end -}}

	for i, v := range a {
		{{if hasPrefix .Of "float" -}}
		// for floats we don't bother checking the incorrect stuff
		if v != correct[i] && i != 0 {
		{{else -}}
		if v != correct[i] {
		{{end -}}
			t.Errorf("{{.OpName}} is incorrect. Expected %v. Got %v", correct[i], v)
			break
		}
	}

	// idiotsville 1
	if err := a.{{.OpName}}({{if eq .Name "f64s"}}float32(2){{else}}2.0{{end}}); err == nil{
		t.Error("Expected an error when performing {{.OpName}} on a differing type")
	}
}
`

const numberVVTestHeaderRaw = `func prep{{.Name}}Test() ({{.Name}}, {{.Name}}, {{.Name}}Dummy, {{.Of}}){
	a := {{.Name}}{ {{.TestData0}} }
	b := {{.Name}}{ {{.TestData1}} }
	c := {{.Name}}Dummy{ {{.TestData1}} }
	return a, b, c, 2
}
`

var (
	binOpTmpl         *template.Template
	binOpTestTmpl     *template.Template
	vecScalarOpTmpl   *template.Template
	vecScalarTestTmpl *template.Template

	vvBinOpTestHeaderTmpl *template.Template
)

func init() {
	binOpTmpl = template.Must(template.New("BinOp").Funcs(funcMap).Parse(binOpRaw))
	binOpTestTmpl = template.Must(template.New("BinOpTest").Funcs(funcMap).Parse(binOpTestRaw))

	vecScalarOpTmpl = template.Must(template.New("vecScalarOp").Funcs(funcMap).Parse(vecScalarOpRaw))
	vecScalarTestTmpl = template.Must(template.New("vecScalarTest").Funcs(funcMap).Parse(vecScalarTestRaw))

	vvBinOpTestHeaderTmpl = template.Must(template.New("vvBinOpTestHeader").Parse(numberVVTestHeaderRaw))
}

func generateNumbers(f io.Writer, m []ArrayType) {
	generateNumbersOpsOnly(f, m)
}

func generateNumbersOpsOnly(f io.Writer, m []ArrayType) {
	// generate V-V bin ops
	for _, bo := range binOps {
		fmt.Fprintf(f, "/* %s */\n\n", bo.OpName)
		for _, v := range m {
			if v.isNumber {
				op := BinOp{v, bo.OpName, bo.OpSymb, bo.IsFunc}
				binOpTmpl.Execute(f, op)
				fmt.Fprintln(f, "\n")
			}
		}
		fmt.Fprintln(f, "\n")
	}

	// generate V-S and S-V bin ops
	for _, bo := range vecscalarOps {
		fmt.Fprintf(f, "/* %s */\n\n", bo.OpName)
		for _, v := range m {
			if v.isNumber {
				op := BinOp{v, bo.OpName, bo.OpSymb, bo.IsFunc}
				vecScalarOpTmpl.Execute(f, op)
				fmt.Fprintf(f, "\n")
			}
		}
		fmt.Fprintf(f, "\n")
	}

	log.Println("NOTE: Manually fix Div for non-float types")

}

func generateNumbersTests(f io.Writer, m []ArrayType) {
	// write headers/prep functions
	for _, v := range m {
		if v.isNumber {
			vvBinOpTestHeaderTmpl.Execute(f, v)
			fmt.Fprintf(f, "\n")
		}
	}

	for _, bo := range binOps {
		fmt.Fprintf(f, "/* %s */\n\n", bo.OpName)
		for _, v := range m {
			if v.isNumber {
				op := BinOp{v, bo.OpName, bo.OpSymb, bo.IsFunc}
				binOpTestTmpl.Execute(f, op)
				fmt.Fprintf(f, "\n")
			}
		}
		fmt.Fprintf(f, "\n")
	}

	for _, bo := range vecscalarOps {
		fmt.Fprintf(f, "/* %s */\n\n", bo.OpName)
		for _, v := range m {
			if v.isNumber {
				op := BinOp{v, bo.OpName, bo.OpSymb, bo.IsFunc}
				vecScalarTestTmpl.Execute(f, op)
				fmt.Fprintf(f, "\n")
			}
		}
		fmt.Fprintf(f, "\n")
	}

}
