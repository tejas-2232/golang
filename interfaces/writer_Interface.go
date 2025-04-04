package main

import (
	"fmt"
)

// define person struct
type Person struct {
	Name string
}

// write String function
// p is called receiver variable:
//   - it is similar to "this" in other prog. languages like java, python

// we are calling String() method on "Person" instance "p"

func (p Person) String() string {

	return "hello- my name is " + p.Name // accessing Person struct using p
}

// writeMessage func creation
func writeMessage(w fmt.Stringer) {
	fmt.Println(w.String())
}

// main func
func main() {

	p := Person{Name: "THOR"}
	writeMessage(p)

}
