package main

import "fmt"

func main() {
	fmt.Println("******************************************************")
	fmt.Println("Pointers in go\n")

	s := 10
	t := &s

	fmt.Println("Value at s= ", s)
	fmt.Println("Value at t=&s (memory address)= ", t)
	fmt.Println("actual value at *t = ", *t)
	fmt.Println("******************************************************")

}
