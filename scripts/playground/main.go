package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/shoriwe/gplasma"
)

type Output struct {
	target *js.Object
}

func (o *Output) Write(p []byte) (n int, err error) {
	// split in multiple lines
	var toWrite []byte
	numberOfCharacterPerLine := js.Global.Get("document").Get("body").Get("clientWidth").Int() * 100 / 1904
	currentLineLength := 0
	for _, character := range p {
		switch character {
		case '\n':
			toWrite = append(toWrite, character)
			currentLineLength = 0
		default:
			// How much is too much
			if currentLineLength == numberOfCharacterPerLine {
				toWrite = append(toWrite, '\n', character)
				currentLineLength = 1
			} else {
				currentLineLength += 1
				toWrite = append(toWrite, character)
			}
		}
	}
	o.target.Set("textContent", o.target.Get("textContent").String()+string(toWrite))
	return len(toWrite), nil
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
