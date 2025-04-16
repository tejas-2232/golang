package main

import (
	"fmt"
	"net/http"
)

// responses with hello there when root URL is accessed
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello There!")

}

func main() {
	http.HandleFunc("/", handler) // used to route incoming requests to our handler func & start server with http.ListenAndServe
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
