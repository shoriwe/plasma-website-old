# Import system

Plasma has a built-in **`Feature`** to import external code from other scripts or for dynamically use external features.

## In Plasma usage

```ruby
# Use `require(MODULE_NAME)` when you want to import plasma or built-in modules
my_module = require("my_module")

# Use `import(SCRIPT_NAME)` when you want to import a script
my_script = import("my-script.pm")

# Use `open_resource(RELATIVE_FILE_PATH)` when you want to have a file handler for a relative file of a module
index_html_template = open_resource("my_module/templates/index.template")

# Use `get_resource_path(RELATIVE_FILE_PATH)` when you want the full file path of a resource in a module
index_html_template_path = get_resource_path("my_module/templates/index.template")  
```

## Loading the feature

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

## Writing your own modules

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

func myModuleLoader() gplasma.ObjectLoader {
	return func(context *gplasma.Context, p *gplasma.Plasma) *gplasma.Value {
		return p.NewInteger(context, false, 0)
	}
}

// Return a valid ModuleInformation to load in the import system
func myModule() importlib.ModuleInformation {
	return importlib.ModuleInformaiton{
		Name:   "my_module",
		Loader: myModuleLoader(),
	}
}

func main() {
	virtualMachine := gplasma.NewVirtualMachine()
	currentDir, err := os.Getwd()
	if err != nil {
		currentDir = "."
	}
	importSystem := importlib.NewImporter()
	// Load Default modules to use with the VM
	importSystem.LoadModule(myModule())
	//
	virtualMachine.LoadFeature(
		importSystem.Result( // Requesting the import system factory the final feature to load in the virtual machine
			importlib.NewRealFileSystem(sitePackagesPath),
			importlib.NewRealFileSystem(currentDir),
		),
	)
}
```

## Using a custom file system search engine

In some scenarios you will need to import existing `Plasma` source code.

For this to succeed you will need to write a new type that implements **`importlib.FileSystem`**.

### Import order

- `import` will always ask only to the provided `pwd` **`importlib.FileSystem`**.

- The loading order when calling `require` is:

1. It searches for a built-in module that has that name
2. If no built-in module was found, it asks the provided `sitePackages` **`importlib.FileSystem`** for a plasma module
   with that name.
