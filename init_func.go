/* init() function is just like the main function, does not take any argument nor return anything.
This function is present in every package and this function is called when the package is initialized.
This function is declared implicitly, so you cannot reference it from anywhere and you are allowed to create multiple init()
function in the same program and they execute in the order they are created.
You are allowed to create init() function anywhere in the program and they are called in lexical file name order (Alphabetical Order)

*/

package main

import "fmt"

func init() {
	fmt.Println("Welcome to init() function")
}

func init() {
	fmt.Println("Hello, init() function")
}

func main() {
	fmt.Println("welcome to main() function")
}
