package main

import "fmt"

// func main() {
// 	fmt.Println("Array:")

// 	var arr1 [3]int = [3]int{1, 2, 3}

// 	fmt.Println("Array=", arr1)
// }

//copying array into another
// both array should have same length and type

var source = [5]int{10, 20, 30, 40, 50}

func main() {

	var destination [5]int

	for i := 0; i < len(source); i++ {
		destination[i] = source[i]
	}

	fmt.Println("source array:", source)
	fmt.Println("destination array:", destination)

}
