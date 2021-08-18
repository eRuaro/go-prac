package main

import "log"

func main() {
	// map
	myMap := make(map[string]string)

	myMap["Neil"] = "Golang"
	myMap["Flutter"] = "Dart"

	// Will print "Golang Dart" and not "Golang Dart null"
	log.Println(myMap["Neil"], myMap["Flutter"], myMap["Golang"])
}
