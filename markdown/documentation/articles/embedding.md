# Embedding

```go
package main

import (
	"fmt"
	"github.com/shoriwe/gplasma"
	"github.com/shoriwe/gplasma/pkg/std/features/importlib"
	"os"
)

var (
	files                   []string
	virtualMachine          *gplasma.VirtualMachine
	sitePackagesPath        = "site-packages"
)

// Setup the vm based on the options
func setupVM() {
	virtualMachine = gplasma.NewVirtualMachine()
	currentDir, err := os.Getwd()
	if err != nil {
		currentDir = "."
	}
	importSystem := importlib.NewImporter()
	// Load Default modules to use with the VM
	importSystem.LoadModule(regex.Regex)
	//
	virtualMachine.LoadFeature(
		importSystem.Result(
			importlib.NewRealFileSystem(sitePackagesPath),
			importlib.NewRealFileSystem(currentDir),
		),
	)
}

func program() {
	setupVM()
	for _, filePath := range files {
		fileHandler, openError := os.Open(filePath)
		if openError != nil {
			_, _ = fmt.Fprintf(os.Stderr, openError.Error())
			os.Exit(1)
		}
		content, readingError := io.ReadAll(fileHandler)
		if readingError != nil {
			_, _ = fmt.Fprintf(os.Stderr, readingError.Error())
			os.Exit(1)
		}
		result, success := virtualMachine.ExecuteMain(string(content))
		if !success {
			_, _ = fmt.Fprintf(os.Stderr, "[%s] %s: %s\n", color.RedString("-"), result.TypeName(), result.String)
			os.Exit(1)
		}
	}
}
```