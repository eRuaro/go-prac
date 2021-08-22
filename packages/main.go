package main

import (
	"github.com/eruaro/mypackage/helper"
	"log"
)

func main() {
	log.Println("Hey")

	var myVar helper.SomeType
	myVar.TypeName = "Some Name"
	myVar.TypeNumber = 5
	log.Println(myVar)

	var myVar2 = helper.SomeType{
		TypeName:   "Hey",
		TypeNumber: 5,
	}

	log.Println(myVar2)
}
