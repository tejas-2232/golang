// an interface is a type that lists methods without providing their code.
// You can’t create an instance of an interface directly, but you can make a variable of the
// interface type to store any value that has the needed methods.

// ## EXAMPLE ##
// type Shape interface {
//     Area() float64
//     Perimeter() float64
// }

// Syntax
//
//	type InterfaceName interface {
//	    Method1() returnType
//	    Method2() returnType
//	}
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length, width float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

// Main function to demonstrate the interface
func main() {
	var s Shape

	s = Circle{radius: 5}
	fmt.Println("C Area:", s.Area())
	fmt.Println("C Perimeter:", s.Perimeter())

	s = Rectangle{length: 4, width: 3}
	fmt.Println("R Area:", s.Area())
	fmt.Println("R Perimeter:", s.Perimeter())
}
