// maps2
// Make me compile!
//
// I AM NOT DONE
package main

import "fmt"

func main() {
	m := map[string]int{
		"John": 22,
		"Ana":  30,
	}

	a, b := m["Ana"]
	fmt.Printf("%#v and %#v\n", a, b)
	fmt.Printf("John is %d and Ana is %d", m["John"], m["Ana"])
}
