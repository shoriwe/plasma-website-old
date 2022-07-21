//go:build js && wasm

package main

import (
	"bytes"
	"sync"
	"syscall/js"
)

const runtimeTerminalId = "runtime-output"

var output *Output

type Output struct {
	terminal js.Value
	mutex    *sync.Mutex
}

func (o *Output) Clear() {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	o.terminal.Set("innerHTML", "")
}

func (o *Output) Write(chunk []byte) (int, error) {
	o.mutex.Lock()
	defer o.mutex.Unlock()
	currentContents := o.terminal.Get("innerHTML").String()
	newContents := make([]byte, 0, len(chunk)+len(currentContents))
	newContents = append(newContents, currentContents...)
	newContents = append(newContents, chunk...)
	newContents = bytes.ReplaceAll(newContents, []byte("\n"), []byte("<br>"))
	o.terminal.Set("innerHTML", string(newContents))
	return len(chunk), nil
}
