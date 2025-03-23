package main

import "fmt"

func main() {
	fmt.Println("******************************************************")
	fmt.Println("Maps in Go \n")

	r := map[string]int{"Alice": 20, "Bob": 25, "David": 30, "Ron": 22}
	fmt.Println("Map stored in r =>", r, "\n")

	fmt.Println("Alice= ", r["Alice"])
	fmt.Println("David= ", r["David"])
	fmt.Println("Bob= ", r["Bob"])
	fmt.Println("Ron= ", r["Ron"])
	fmt.Println("******************************************************")

}
