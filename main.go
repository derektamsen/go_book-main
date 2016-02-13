package main

import (
	"fmt"
)

type Add func(a int, b int) int

func main() {
	fmt.Println(process(func(a int, b int) int {
		return a + b
	}))
}

func process(adder Add) int {
	return adder(1, 2)
}
