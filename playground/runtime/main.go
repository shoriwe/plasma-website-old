//go:build js && wasm

package main

import (
	"github.com/shoriwe/gplasma/pkg/vm"
	"sync"
	"syscall/js"
)

const runCodeFunctionName = "runCode"

func init() {
	output = &Output{
		terminal: js.Global().Get(runtimeTerminalId),
		mutex:    &sync.Mutex{},
	}
	editor = js.Global().Get(editorId)
	plasma = vm.NewVM(nil, output, output)

	js.Global().Set(runCodeFunctionName, runCode())
}

func main() {
	wait()
}

func wait() {
	<-make(chan struct{})
}
