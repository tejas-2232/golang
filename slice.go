/* Slices - Composite type

 Dynamic size: slices are similar to arrays but their size can change. They are more flexible
and are used more often than arrays

p:= []int{1,2,3,4,5,6}  //slice of integers
*/

package main

import "fmt"

func main() {

	fmt.Println("Slice:")
	dyn_arr := []int{1, 2, 3, 4, 5, 6} // can't add float and string

	fmt.Println("slice is=>", dyn_arr)

}
