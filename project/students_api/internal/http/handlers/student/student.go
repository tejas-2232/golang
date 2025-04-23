package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/tejas-2232/students_api/internal/types"
	"github.com/tejas-2232/students_api/internal/utils/response"
)

func New() http.HandlerFunc {
	// return data
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student //see this struct in types/types.go

		//r in r.Body can implement Decoder interface
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) { // target error is io.EOF
			response.WriteJson(w, http.StatusBadRequest, err.Error()) // code 400
			return
		}

		// serialize
		// send json data to struct
		slog.Info("creating a student")

		// w.Write([]byte("Welcome to students API")) //converting string to byte slice & passing to write func

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"}) // code 201- created
	}
}
