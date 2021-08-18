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
	// short hand for var anotherWord string = "Another Word"
	anotherWord := "Another Word"
	fmt.Println(word)
	fmt.Println(hello)
	log.Println(word)
	log.Println(anotherWord)
}

//functions
func saySomething(s string) string {
	return s
}