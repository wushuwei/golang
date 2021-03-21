package main

import (
	"fmt"

	"github.com/wushuwei/golang/test/functions"
)

func main() {
	// functions.Data()
	_, _, e := functions.Functions(1, 0)
	if e == nil {
		fmt.Print("ok")
	} else {
		fmt.Print(e)
	}
}
