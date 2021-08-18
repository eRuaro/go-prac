package main

import (
	"log"
)

// Start with capital so that it would be "public"
type Person struct {
	// Capital for "public" viewing
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

	log.Println(neil.Name)
	log.Println(neil.Age)
	log.Println(neil)
}

// visible to only this package
func saySomething(word string) {
	log.Println("Saying:", word)
}
