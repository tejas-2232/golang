package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {

	s := []int{234, 43, 432, 5, 65, 653, 3}

	sort.Ints(s)
	fmt.Println("Sorted slice= ", s)

	//find index
	fmt.Println("index value = ", strings.Index("HelloThere", "ks"))

	// finding time

	fmt.Println("Time: ", time.Now().Unix())
}
