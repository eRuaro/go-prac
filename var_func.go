package main

import (
	"fmt"
	"log"
)

func main() {
	// Printing hello world
	fmt.Println("Hello, World!")

	//variables
	var word string = saySomething("hey")
	var hello string = "hello"
	fmt.Println(word)
	fmt.Println(hello)
	log.Println(word)
}

//functions
func saySomething(s string) string {
	return s
}