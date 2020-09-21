# Hex Color Library for Go

![Test](https://github.com/6543/go-hexcolor/workflows/Test/badge.svg?event=push) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![GoDoc](https://godoc.org/github.com/6543/go-hexcolor?status.svg)](https://godoc.org/github.com/6543/go-hexcolor) [![Go Report Card](https://goreportcard.com/badge/github.com/6543/go-hexcolor)](https://goreportcard.com/report/github.com/6543/go-hexcolor)

## Features:
 * Normalize/Parse hex color
 * convert from/to color.RGBA

## ToDo
* [ ] Predefine standard colors (red, blue, ... based on CSS colors)

## Usage

Download: `go get -u github.com/6543/go-hexcolor`

Example:
```go
import (
	"fmt"
	"github.com/6543/go-hexcolor"
)

func main() {
	c, err := hexcolor.NewHexColor("#adf")
	if err != nil {
		panic(err)
	}
	fmt.Println(c.ToString())
}
```

## Contribution
Fork repository, clone it, make changes, push to new branch and submit a pull request.
