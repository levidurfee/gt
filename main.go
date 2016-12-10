package main

import (
	"fmt"
	"strconv"
	"github.com/levidurfee/stringutil"
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

type Petter interface {
	GivePets(p Pet) string
}

func (p *Person) GivePets(pet Pet) {
	fmt.Printf(p.Name + " gives " + pet.PetName + " 100 pets.. aww.. good pet!\n")
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
	p.GivePets(p.Pet)

	gone("bye!")
	fmt.Printf(stringutil.Reverse("Oh helloo!!!"))
}
