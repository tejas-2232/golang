// Type assertion allows extracting the underlying type of an interface.
// SYNTAX:
//       value := interfaceVariable.(ConcreteType)

package main

import "fmt"

func printArea(s interface{}) {
	value, ok := s.(Shape)
	if ok {
		fmt.Println("Area:", value.Area())
	} else {
		fmt.Println("Not a shape interface")
	}
}

type Shape interface {
	Area() float64
}

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func main() {

	sq := Square{side: 4}
	printArea(sq)
}
