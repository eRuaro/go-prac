package main

import "fmt"

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
}
