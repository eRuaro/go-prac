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

	log.Println(returnMult(5, 10))
	log.Println(returnMult3(5, 10, "Hey"))
}

//functions
func saySomething(s string) string {
	return s
}

//returning multiple things
func returnMult(a, b int) (int, int) {
	return a, b
}

func returnMult3(a int, b int, c string) (int, int, string) {
	return a, b, c
}