package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"First Name"` // must match that of receiving JSON
	LastName  string `json:"Last Name"`
}

func main() {
	// raed json
	myJson := `
[
			{
				"First Name": "Clark",
				"Last Name": "Kent"
			},
			{
				"First Name": "Bruce",
				"Last Name": "Wayne"
			}
		]`

	var unmarshalled []Person // json returns array
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}

	log.Printf("Unmarshalled: %v", unmarshalled)

	// write json from struct
	var mySlice []Person

	var m1 Person = Person{
		FirstName: "Rawo",
		LastName:  "Rawo",
	}

	mySlice = append(mySlice, m1)

	// convert to json
	// marshalindent for better formatting like json extensions in chome
	newJson, err := json.Marshal(mySlice)
	newJson2, err2 := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		log.Println("Error marshalling", err)
	}

	if err2 != nil {
		log.Println("Error marshalling", err2)
	}

	fmt.Println(string(newJson)) // not putting it in string will put in numbers
	fmt.Println(string(newJson2))
}
