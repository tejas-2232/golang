/*
An anonymous function is a function that doesn’t have a name.
It is useful when you want to create an inline function.
In Go, an anonymous function can also form a closure. An anonymous function is also known as a function literal
*/

// package main
// import "fmt"
// func main() {
// 	// Anonymous function
// 	func() {
// 		fmt.Println("The american revolution")
// 	}()
// }

/*
Assigning to a Variable
You can assign an anonymous function to a variable. This variable can then be called like a regular function.*/

// package main
// import ("fmt")
// func main() {
// 	value := func() {
// 		fmt.Println("Golang is the new future ")
// 	}

// 	value() // calling var value as function
// }

/*Passing Arguments
Anonymous functions can accept arguments.*/

package main

import "fmt"

func main() {
	func(ele string) {
		fmt.Println(ele)
	}("French revolution")
}
