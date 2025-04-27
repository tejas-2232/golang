package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tejas-2232/students_api/internal/types"
	"github.com/tejas-2232/students_api/internal/utils/response"
)

func New() http.HandlerFunc {
	// return data
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("creating a student")
		var student types.Student //see this struct in types/types.go

		//r in r.Body can implement Decoder interface
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) { // target error is io.EOF
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body"))) // code 400
			return
		}
		// IF error is other than empty body

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// request validation
		if err := validator.New().Struct(student); err != nil {
			validateErrs := err.(validator.ValidationErrors) //type casting the error, type is ValidationErrors

			response.WriteJson(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}
		// serialize
		// send json data to struct
		// w.Write([]byte("Welcome to students API")) //converting string to byte slice & passing to write func

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"}) // code 201- created
	}
}
