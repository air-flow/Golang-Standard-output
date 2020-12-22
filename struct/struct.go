package main

import "fmt"

// Person test
type Person struct {
	firstName string
	age       int
}

func (p *Person) set() {
	p.age = 1
	p.firstName = "test"
	// return p
}

func main() {
	bob := new(Person)
	bob.set()
	fmt.Println(bob.age)
	fmt.Print(bob)
}
