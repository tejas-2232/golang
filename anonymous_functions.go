/*
An anonymous function is a function that doesn’t have a name.
It is useful when you want to create an inline function.
In Go, an anonymous function can also form a closure. An anonymous function is also known as a function literal
*/
//*********************************************************
// RUN EACH BLOCK AS SEPARATE CODE BY UNCOMMENTING IT   -  T H A N K S
//*********************************************************

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

// package main
// import "fmt"
// func main() {
// 	func(ele string) {
// 		fmt.Println(ele)
// 	}("French revolution")
// }

/*Passing as Arguments
You can also pas an anonymous func as argument to another function
*/

// package main

// import "fmt"

// func galaxy(i func(p, q string) string) {
// 	fmt.Println(i("Golden ", "Hour "))
// }

// // Think of i as a placeholder for a function.
// // Just like we pass integers or strings as arguments, we are passing a function here.

// func main() {

// 	value := func(p, q string) string {
// 		return p + q + "at beach"
// 	}
// 	galaxy(value) // The galaxy function is called with value as its argument
// }

// Example: Multiple Function Parameters

package main

import (
	"fmt"
	"strings"
)

func process(a, b string, f1 func(string, string) string, f2 func(string) string) {

	result := f1(a, b)        //call first function
	finalresult := f2(result) // call second function
	fmt.Println(finalresult)
}

func main() {

	// firts anonymous functions concats 2 strings
	concat := func(p, q string) string {

		return p + " & " + q
	}

	toUpper := func(s string) string { //convert string to upper case
		return strings.ToUpper(s)
	}

	//call process with 2 functions

	process("Golder Hour ", "Beach Sunset", concat, toUpper)

}
