package response

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	// we are gonna pass write response object
	// we wrote any for datatype of data like= "data any", because we don't know what type of data we'll  be receiving
	// OR we can write empty interface like = "data interface{}"
	// return type of func is  = error

	w.Header().Set("Content-Type", "application/json") // to send the data
	w.WriteHeader(status)

	// for incoming data, to put in in struct we decoded it.
	//we will need to encode here as we are converting struct info to json, so

	return json.NewEncoder(w).Encode(data) // Encode method returns error -

}
