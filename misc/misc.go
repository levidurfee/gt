package misc

import (
	"fmt"
	"github.com/levidurfee/gt/prime"
	"github.com/levidurfee/stringutil"
	"strconv"
)

type Person struct {
	Id   int
	Name string
	Age  int
	Pet  Pet
}

type Pet struct {
	PetName string
	PetType string
	PetAge  int
}

type Petter interface {
	GivePets(p Pet) string
}

func (p *Person) GivePets(pet Pet) string {
	return p.Name + " gives " + pet.PetName + " 100 pets.. aww.. good pet!"
}

func (p *Person) getPet() {
	dogAge := strconv.Itoa(p.Pet.PetAge)
	dogYears := strconv.Itoa(p.Pet.PetAge * 7)
	fmt.Println("Hi, I am " + p.Name + " and I have a pet named " + p.Pet.PetName + ", he is " + dogAge +
		" in human years. That makes him " + dogYears + " in dog years!")
}

func gone(s string) {
	fmt.Println("Gone: " + s)
}

/*func main() {
	p := Person{
		Id:   1,
		Name: "Levi",
		Age:  32,
		Pet: Pet{
			PetName: "Corky",
			PetType: "Mixed",
			PetAge:  9,
		},
	}

	p.getPet()
	fmt.Println(p.GivePets(p.Pet))

	gone("bye!")
	fmt.Println(stringutil.Reverse("Oh helloo!!!"))

	primes := prime.Sieve(100)
	for _, p := range primes {
		fmt.Printf(strconv.Itoa(p) + ", ")
	}
}
*/
