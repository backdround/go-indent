[![Go Reference](https://img.shields.io/badge/go-reference-%2300ADD8?style=flat-square)](https://pkg.go.dev/github.com/backdround/go-indent)
[![Tests](https://img.shields.io/github/workflow/status/backdround/go-indent/tests?label=tests&style=flat-square)](https://github.com/backdround/go-indent/actions)
[![Codecov](https://img.shields.io/codecov/c/github/backdround/go-indent?style=flat-square)](https://app.codecov.io/gh/backdround/go-indent/)
[![Go Report](https://goreportcard.com/badge/github.com/backdround/go-indent?style=flat-square)](https://goreportcard.com/report/github.com/backdround/go-indent)

# Indent
Indent implements function that allows to indent multiline strings

### Installation
```bash
go get github.com/backdround/go-indent
```

### Example

```go
package main

import (
	"fmt"
	"github.com/backdround/go-indent"
)

func main() {
	message := fmt.Sprint("some\nmultiline\nmessage")
	indentedMessage := indent.Indent(message, "  ", 1)
	fmt.Println(indentedMessage)
}
```

output:
```
  some
  multiline
  message
```
