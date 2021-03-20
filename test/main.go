package main

import (
	"fmt"
)

func main() {
	str := "world"
	_, err := fmt.Printf("Type: %T\n", str)
	if err != nil {
	}
}
