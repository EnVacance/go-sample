package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

func main() {

	// Init person using struct
	person1 := Person{
		firstName: "Samantha",
		lastName:  "Smith",
		city:      "Boston",
		gender:    "f",
		age:       20,
	}

	fmt.Println(person1)
	person1.age++
	fmt.Println(person1)

	person2 := person1
	fmt.Println(person2)
	person2.age = 300
	fmt.Println(person1)
	fmt.Println(person2)

	fmt.Println(person1.greet())

	greet(person1)
}
