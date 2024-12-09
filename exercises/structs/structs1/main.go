// structs1
// Make me compile!
//
// I AM NOT DONE
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	person := Person{name: "Beck", age: 22}
	fmt.Printf("Person %s and age %d", person.name, person.age)
}
