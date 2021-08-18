package main

import (
	"fmt"
)

func main() {
	var a int = 1
	// takes in memory address of a
	var b *int = &a
	fmt.Println(a, *b)
	fmt.Println(b, &a)

	// change value of a
	a = 2
	// b is also changed due to having same memory address of a
	fmt.Println(a, *b)
	fmt.Println(b, &a)
}