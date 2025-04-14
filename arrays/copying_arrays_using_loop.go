//copying array into another

package main

import "fmt"

// both array should have same length and type
var source = [5]int{10, 20, 30, 40, 50}

func main() {

	var destination [5]int

	for i := 0; i < len(source); i++ {
		destination[i] = source[i]
	}
	fmt.Println("Source array:", source)
	fmt.Println("destination array:", destination)

}
