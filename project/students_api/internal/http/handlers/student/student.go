package student

import "net/http"

func New() http.HandlerFunc {
	// return data
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students API")) //converting string to byte slice & passing to write func
	}
}
