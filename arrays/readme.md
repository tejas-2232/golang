# copying array into another

## method 1: Using a Loop to Copy an Array

Go doesn’t provide a built-in copy() function for arrays, so the most common way to copy an array is by iterating through each element and assigning it manually

__syntax:__

```go
for i:=0; i< len(source);i++{
    destination[i] = source[i]
}
```

__Example:__

```go
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
```