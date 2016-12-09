package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	Id int
	Name string
	Age int
	Pet Pet
}

type Pet struct {
	PetName string
	PetType string
	PetAge int
}

func (p *Person) getPet() {
	dogAge := strconv.Itoa(p.Pet.PetAge)
	dogYears := strconv.Itoa(p.Pet.PetAge * 7)
	fmt.Printf("Hi, I am " + p.Name + " and I have a pet named " + p.Pet.PetName + ", he is " + dogAge + " in human years. That makes him " + dogYears + " in dog years!\n")
}

func gone(s string) {
	fmt.Printf(s + "\n")
}

func main() {
	p := Person {
		Id: 1,
		Name: "Levi",
		Age: 32,
		Pet: Pet {
			PetName: "Corky",
			PetType: "Mixed",
			PetAge: 9,
		},
	}

	p.getPet()

	gone("bye!")
}
