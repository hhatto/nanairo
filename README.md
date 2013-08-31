# nanairo
nanairo is a Go library that make the colorized output of terminal applications.

[![Build Status](https://travis-ci.org/hhatto/nanairo.png?branch=master)](https://travis-ci.org/hhatto/nanairo)

## Installation

```
go get "github.com/hhatto/nanairo"
```

## Usage
```go
package main
import (
    "fmt"
    nanairo "github.com/hhatto/nanairo"
)

func main() {
    fmt.Println(nanairo.FgColor("#c93", "Hello World"))
}
```

## License
MIT
