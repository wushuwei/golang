package functions

import (
	"errors"
	"fmt"
)

func HelloWorld() {
	// str := "world"
	// fmt.Printf("Type: %T\n", str)

	// const y = "hello"
	// fmt.Printf("Type: %T\n", y)

	// x := [3]int{10, 20, 30}
	// fmt.Printf("Type: %T\n", x)

	// z := [...]int{10, 20, 30}
	// fmt.Printf("Type: %T\n", z)

	m := make([]int, 0, 10)
	fmt.Printf("Type: %T\n", m)
	m = append(m, 5, 6, 7, 8)
	fmt.Printf("Type: %T\n", m)
	fmt.Printf("Values: %v\n", m)

	t := map[string][]string{
		"Orcas":   {"Fred", "Ralph", "Bijou"},
		"Lions":   {"Sarah", "Peter", "Billie"},
		"Kittens": {"Waldo", "Raul", "Ze"},
	}
	fmt.Printf("Type: %T\n", t)
	fmt.Printf("Values: %v\n", t)

	for _, v := range t {
		fmt.Printf("v only: %v\t", v)
	}

	if _, ok := t["Lions"]; ok {
		fmt.Println("found!")
	}

}

func Data() {
	type p struct {
		name string
		age  int
		pet  string
	}
	bob := p{}
	fmt.Printf("Type: %T\n", bob)
	fmt.Printf("Value: %v\n", bob)

	julia := p{"Julia", 40, "cat"}
	fmt.Printf("Type: %T\n", julia)
	fmt.Printf("Value: %v\n", julia)
}

func Functions(i int, j int) (int, int, error) {
	if i == 0 {
		return 0, 0, errors.New("it is an error!")
	}
	return 1, 1, nil
}
