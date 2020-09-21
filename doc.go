// Copyright 2020 6543. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*

Package hexcolor is a simple hexcolor normalizer

Installation

    go get -u github.com/6543/go-hexcolor

Example

    ```go
    package main

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

*/
package hexcolor
