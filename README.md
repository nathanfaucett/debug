Debugger.go
=====

Debugger.go is simple debug function that returns a struct for debugging
if env's GO_ENVIRONMENT_NAME = Dev or Development

##Example
```
package main

import (
	"github.com/nathanfaucett/debugger"
	"errors"
)

var debug = debugger.Debug("WebApp")

func main() {
	debug.Log("Main Function Called")
	debug.Warning("Main Function Called")
	debug.Error("Main Function Called")
	debug.Error(errors.New("Main Function Called"))
}

// output
// WebApp: Main Function Called +0ms
// WebApp Warning: Main Function Called +0ms
// WebApp Error: Main Function Called +0ms
// WebApp Error: Main Function Called +0ms

```