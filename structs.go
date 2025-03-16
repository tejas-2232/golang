/*
- create custom data type using struct
- give name as person and add Name, Age in that struct
- print the struct
*/

package main

import "fmt"

func main() {

	fmt.Println("Struct:")

	type Person struct {
		Name string
		Age  int
	}

	var x Person = Person{"Andrew", 29}

	fmt.Println("x=", x)
	fmt.Println("Name=", x.Name)
	fmt.Println("Age=", x.Age)

}
