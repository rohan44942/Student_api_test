package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rohan44942/student-api/internal/storage"
	"github.com/rohan44942/student-api/internal/types"
	"github.com/rohan44942/student-api/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}
		// w.Write([]byte("Hello, bro phla go ka api hit krne ke liye congratulations dude !!!"))
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		// validate the request
		if err := validator.New().Struct(student); err != nil {
			validateErros := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusForbidden, response.ValidationError(validateErros))
			return
		}
		// now store the student
		lastId, err := storage.CreateStudent(student.Name, student.Email, student.Age)
		slog.Info("user created successfuly ", slog.String("user Id", fmt.Sprint(lastId)))
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusCreated, map[string]any{"id": lastId,
			"succes": "ok"})
	}
}

func StudentInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		// Use id to fetch and return student data
		if id == "6" {
			w.Write([]byte("hello , i am rohan with id 6 as you passed"))
		}
	}
}
