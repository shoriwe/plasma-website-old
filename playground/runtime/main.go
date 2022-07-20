//go:build wasm

package main

import (
	"bytes"
	"fmt"
	"github.com/shoriwe/gplasma/pkg/vm"
	"syscall/js"
)

var (
	plasma                *vm.Plasma
	errorChannel          chan error
	stopChannel           chan struct{}
	output                *Output
	runtimeOutput, editor js.Value
)

type Output struct {
}

func (output *Output) Write(c []byte) (int, error) {
	c = bytes.ReplaceAll(c, []byte("\n"), []byte("<br>"))
	runtimeOutput.Set("innerHTML",
		runtimeOutput.Get("innerHTML").String()+string(c),
	)
	return len(c), nil
}

func onError(err error) []byte {
	return []byte(fmt.Sprintf(
		`<br><p style="color: #e74c3c;">%s</p>`,
		err.Error(),
	))
}

func init() {
	output = &Output{}
	plasma = vm.NewVM(nil, output, output)
	editor = js.Global().Get("editor")
	runtimeOutput = js.Global().Get("document").Call("getElementById", "runtime-output")
}

func runCode() {
	runtimeOutput.Set("innerText", "")
	code := editor.Call("getValue").String()
	if stopChannel != nil {
		stopChannel <- struct{}{}
	}
	_, errorChannel, stopChannel = plasma.ExecuteString(code)
	go func(ec chan error) {
		err := <-ec
		if err != nil {
			output.Write(onError(err))
		}
	}(errorChannel)
}
func main() {
	js.Global().Set("runCode",
		js.FuncOf(func(_ js.Value, _ []js.Value) interface{} {
			runCode()
			return nil
		}),
	)
	<-make(chan struct{})
}
