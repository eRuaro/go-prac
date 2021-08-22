package main

import (
	"github.com/eruaro/mychannel/helpers"
	"log"
)

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(10)

	intChan <- randomNumber
}

func main() {
	PrintText("Neil")

	//channels
	intChan := make(chan int)
	defer close(intChan)

	// goroutine
	go CalculateValue(intChan)

	//listen to response
	num := <-intChan
	log.Println(num)
}

// procedural way
func PrintText(s string) {
	log.Println(s)
}
