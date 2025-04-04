package main

import (
	"fmt"
)

// define person struct
type Person struct {
	Name string
}

// write String function
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
