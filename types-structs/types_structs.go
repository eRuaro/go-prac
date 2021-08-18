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

type withFunction struct {
	FirstName string
}

// receiver allows to use functions for structs
func (m *withFunction) printFirstName() string {
	return m.FirstName
}

func main() {
	//word := "Hey"
	//log.Println("Word is,", word)
	//saySomething(word)

	neil := Person{
		Name: "Neil",
		Age:  17,
	}

	log.Println(neil.Name)
	log.Println(neil.Age)
	log.Println(neil)

	myVar := withFunction{
		FirstName: "My Var",
	}

	log.Println("My Var is set to:", myVar.FirstName)
	// accessing a member using a function
	log.Println(myVar.printFirstName())
}

// visible to only this package
//func saySomething(word string) {
//	log.Println("Saying:", word)
//}
