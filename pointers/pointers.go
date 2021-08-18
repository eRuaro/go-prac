package main

import (
	"log"
)

func main() {
	var a int = 1
	// takes in memory address of a
	var b *int = &a
	log.Println(a, *b)
	log.Println(b, &a)

	// change value of a
	a = 2
	// b is also changed due to having same memory address of a
	log.Println(a, *b)
	log.Println(b, &a)

	point(&a, b)
}

func point(x *int, y *int) {
	log.Println(x, y)
}
