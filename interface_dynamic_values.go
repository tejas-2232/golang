// An interface can hold any value, but the actual value and its type are stored dynamically.

package main

import "fmt"

type Shape interface {
	Area() float64
}

func main() {
	var s Shape
	fmt.Println("Value of s:", s)    // nil
	fmt.Printf("Type of s: %T\n", s) // <nil>
}
