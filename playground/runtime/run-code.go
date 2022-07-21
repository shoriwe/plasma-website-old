//go:build js && wasm

package main

import (
	"fmt"
	"github.com/shoriwe/gplasma/pkg/vm"
	"syscall/js"
)

var (
	activeStopChannel chan struct{}
	editor            js.Value
	plasma            *vm.Plasma
)

const (
	editorId       = "editor"
	editorGetValue = "getValue"
)

func getCode() string {
	return editor.Call(editorGetValue).String()
}

func handleError(errorChannel chan error) {
	err := <-errorChannel
	if err != nil {
		data := fmt.Sprintf(`<p style="color: red;">%s</p>`, err.Error())
		_, _ = output.Write([]byte(data))
	}
}

func runCode() js.Func {
	return js.FuncOf(
		func(this js.Value, args []js.Value) any {
			output.Clear()
			if activeStopChannel != nil {
				activeStopChannel <- struct{}{}
			}
			var errorChannel chan error
			_, errorChannel, activeStopChannel = plasma.ExecuteString(getCode())
			go handleError(errorChannel)
			return 0
		},
	)
}
