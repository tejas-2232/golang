/*
In Go language, defer statements delay the execution of the function or method or an anonymous method until the
nearby functions returns. In other words, defer function or method call arguments evaluate instantly,
but they don’t execute until the nearby functions returns. You can create a deferred method, or function, or
 anonymous function by using the defer keyword.
*/
// concept of the defer statement

// package main

// import "fmt"

// func mul(a1, a2 int) int {

// 	res := a1 * a2
// 	fmt.Println("Result= ", res)
// 	return 0
// }

// func show() {
// 	fmt.Println("Hello there")
// }

// func main() {
// 	// Calling mul() function
// 	// Here mul function behaves
// 	// like a normal function
// 	mul(20, 5)

// 	// Calling mul()function Using defer keyword
// 	// Here the mul() function is defer function

// 	defer mul(20, 5)

// 	// Calling show() function
// 	show()

// }

// Example 2

// Go program to illustrate
// multiple defer statements, to illustrate LIFO policy

package main

import "fmt"

func mul(a1, a2 int) int {
	res1 := a1 + a2
	fmt.Println("Result= ", res1)
	return 0
}

func main() {

	fmt.Println("Start")

	defer fmt.Println("End") // 3rd

	defer mul(12, 12) // = 24  // 2nd

	defer mul(10, 30) // = 40  // 1st

}
