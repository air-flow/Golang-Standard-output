package main

import (
	"fmt"
)

// Person test
type Person struct {
	firstName string
	age       int
}

func (p Person) intro(greetings string) string {
	return greetings + " I am " + p.firstName
}

func main() {
	bob := Person{age: 15, firstName: "Sam"}
	fmt.Println(bob.intro("Hello")) //=> Hello I am Bob
}
