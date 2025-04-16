// serving dynamic content
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// rand.Seed(time.Now().UnixNano()) // moving this to main() so that random generator seeded only once
	// leads to better randomness over time

	randomNumber := rand.Intn(100)
	fmt.Fprintf(w, "Random Number: %d", randomNumber)
}
func main() {
	rand.Seed(time.Now().UnixNano()) // Seeding once at the start of the program
	http.HandleFunc("/", handler)
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
