package response

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

type Response struct {
	Status string `json:"status`
	Error  string `json:"error` // "Error" will be displayed as "error"
} // because we want to send a json like that if we get any error

const (
	StatusOK    = "OK"
	StatusError = "Error"
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
	// when we dont include json data in postman while testing it throws "EOF" response
	// it should throw json -
	// ######################################
	// solution- add Response struct and create one generalError method
}

func GeneralError(err error) Response {
	return Response{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErros) Response {
	var errMsgs []string //create slice os string

	for _, err := range errs {

	}
}
