package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method ( value recieved)
func (person Person) greet() string {
	return "\n\nHello, my name is " + person.firstName + " " + person.lastName + "\nI'm " + strconv.Itoa(person.age)
}

// hasBirthday method ( pointer reciever )
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried ( pointer reciever )
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "m" {
		return
	} else {
		person.lastName = spouseLastName
	}
	fmt.Println("Till death shall it be!")
}

func main() {
	person1 := Person{firstName: "Rigena", lastName: "Rein", city: "Boston", gender: "f", age: 25}
	person2 := Person{"Ivan", "Divan", "Brighton", "m", 22}

	fmt.Println(person1.greet())

	person1.age = person2.age
	person1.hasBirthday()
	fmt.Println(person1)

	person1.getMarried(person2.lastName)
	fmt.Println(person1)

}
