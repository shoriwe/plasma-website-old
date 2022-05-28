# Features

To limit the runtime capabilities of gplasma when embedding, the user has the ability to decide which ones he wanted to
have enabled during execution.

## Loading features

Once created an instance of the virtual machine, you can approach **`LoadFeature(feature)`** Go function load the new
feature you want.

```go
package main

import (
	"github.com/shoriwe/gplasma/pkg/std/features/importlib"
	"github.com/shoriwe/gplasma"
	"github.com/shoriwe/gplasma/pkg/std/modules/base64"
	"github.com/shoriwe/gplasma/pkg/std/modules/json"
	"github.com/shoriwe/gplasma/pkg/std/modules/regex"
	"os"
)

func main() {
	virtualMachine := gplasma.NewVirtualMachine()
	currentDir, err := os.Getwd()
	if err != nil {
		currentDir = "."
	}
	importSystem := importlib.NewImporter()
	// Load Default modules to use with the VM
	importSystem.LoadModule(regex.Regex)
	importSystem.LoadModule(json.JSON)
	importSystem.LoadModule(base64.Base64)
	//
	virtualMachine.LoadFeature(
		importSystem.Result( // Requesting the import system factory the final feature to load in the virtual machine
			importlib.NewRealFileSystem(sitePackagesPath),
			importlib.NewRealFileSystem(currentDir),
		),
	)
}
```

## Built-in features

- [Import system](/plasma/documentation/features-import-system.html)

## Writing your own features

A feature is just a hash map of strings and [ObjectLoaders](/plasma/documentation/features-object-loader.html). Where after
executing the **`ObjectLoader`**, it assigns its result to the name specified by the string in the virtual machine
master namespace.

### Example

```go
package main

import (
	"github.com/shoriwe/gplasma/pkg/std/features/importlib"
	"github.com/shoriwe/gplasma"
	"github.com/shoriwe/gplasma/pkg/std/modules/base64"
	"github.com/shoriwe/gplasma/pkg/std/modules/json"
	"github.com/shoriwe/gplasma/pkg/std/modules/regex"
)

// This function will return the feature to load in the VM
func myFeature() gplasma.Feature {
	return map[string]gplasma.ObjectLoader{
		"custom_object": func(context *gplasma.Context, p *gplasma.Plasma) *gplasma.Value {
			return gplasma.NewValue(context, false, gplasma.ValueName, nil, context.PeekSymbolTable())
		},
	}
}

func main() {
	virtualMachine := gplasma.NewVirtualMachine()
	// Loading the feature
	virtualMachine.LoadFeature(
		myFeature(),
	)
}
```