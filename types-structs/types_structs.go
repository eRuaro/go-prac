package main

import (
	"log"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	word := "Hey"
	log.Println("Word is,", word)
	saySomething(word)

	neil := Person{
		Name: "Neil",
		Age:  17,
	}

	log.Println(neil)
}

func saySomething(word string) {
	log.Println("Saying:", word)
}
