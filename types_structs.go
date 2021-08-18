package main

import(
	"log"
)

func main() {
	word := "Hey"
	log.Println("Word is,", word)	
	saySomething(word)
}

func saySomething(word string) {
	log.Println("Saying:", word)
}