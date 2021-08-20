package main

import (
	"fmt"
	"log"
)

func main() {
	for number := 0; number < 20; number++ {
		if number < 5 {
			fmt.Println("Number:", number, " is less than 5")
		} else if number < 10 {
			fmt.Println("Number:", number, " is less than 10")
		} else if number < 15 {
			fmt.Println("Number:", number, " is less than 15")
		} else {
			fmt.Println("Number:", number, " is less than 20")
		}
	}

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Ranging over Slice")
	for _, x := range mySlice {
		log.Println(x)
	}

	myMap := make(map[string]int)
	myMap["Dog"] = 5
	myMap["Cat"] = 10
	myMap["Fish"] = 15
	fmt.Println("Ranging over Map")
	for i, x := range myMap {
		log.Println(i, x)
	}
}
