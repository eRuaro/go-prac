package main

import "log"

//Abstract classes
type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Adam",
		Breed: "German Shepherd",
	}

	PrintInfo(dog)

	gorilla := Gorilla{
		Name:          "Alex",
		Color:         "Silver",
		NumberOfTeeth: 32,
	}

	PrintInfo(gorilla)
}

func (d Dog) Says() string {
	return "woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func (g Gorilla) Says() string {
	return "roar"
}

func (g Gorilla) NumberOfLegs() int {
	return 2
}

func PrintInfo(a Animal) {
	log.Println(a.Says())
	log.Println(a.NumberOfLegs())
}
