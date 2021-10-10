package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/shoriwe/gplasma"
)

type Output struct {
	target *js.Object
}

func (o *Output) Write(p []byte) (n int, err error) {
	o.target.Set("textContent", o.target.Get("textContent").String()+string(p))
	return len(p), nil
}

func NewOutput() *Output {
	return &Output{
		target: js.Global.Get("output"),
	}
}

var (
	virtualMachine *gplasma.VirtualMachine
	output         *Output
)

func init() {
	output = NewOutput()
	virtualMachine = gplasma.NewVirtualMachine()
	virtualMachine.Stderr = output
	virtualMachine.Stdout = output
	virtualMachine.Stdin = nil
}

func runCode() {
	output.target.Set("textContent", "")
	codeContainer := js.Global.Get("code_container")
	contents := codeContainer.Get("value")
	result, executionSuccess := virtualMachine.ExecuteMain(contents.String())
	if executionSuccess {
		return
	}
	errorMessage := result.GetClass(virtualMachine.Plasma).Name + ": " + result.String + "\n"
	_, _ = output.Write([]byte(errorMessage))
}

func main() {
	js.Global.Set("run_code", runCode)
}
