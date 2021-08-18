package main

import "log"

func main() {
	// map -> HashTable
	myMap := make(map[string]string)

	myMap["Neil"] = "Golang"
	myMap["Flutter"] = "Dart"

	// Will print "Golang Dart" and not "Golang Dart null"
	log.Println(myMap["Neil"], myMap["Flutter"], myMap["Golang"])

	// slices -> Dynamic Array
	var myArray []string
	myArray = append(myArray, "Neil")
	myArray = append(myArray, "Golang", "Flutter")

	log.Println(myArray)
	log.Println(myArray[1])

	// slices short hand
	numbers := []int{5, 8}
	numbers = append(numbers, 11)
	log.Println(numbers)
}
