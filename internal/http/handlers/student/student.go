package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

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

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// main logic goes here
		// id from params
		id := r.PathValue("id")
		slog.Info("getting a student by id", slog.String("id", id))
		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		student, err := storage.GetStudentById(intId)

		if err != nil {
			slog.Error("error getting user", slog.String("id", id))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, student)
	}
}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// main logic goes here
		slog.Info("getting all students")
		students, err := storage.GetStudents()
		if err != nil {
			response.WriteJson(w, http.StatusInternalServerError, err)
			return
		}

		response.WriteJson(w, http.StatusOK, students)
	}
}

func UpdateById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// main logic goes here
		// id from params
		id := r.PathValue("id")
		slog.Info("updating a student by id", slog.String("id", id))

		var studentBody types.Student
		err := json.NewDecoder(r.Body).Decode(&studentBody)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		// validate the request
		if err := validator.New().Struct(studentBody); err != nil {
			validateErros := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusForbidden, response.ValidationError(validateErros))
			return
		}

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		// student, err := storage.GetStudentById(intId)
		student, err := storage.UpdateStudentById(intId, studentBody.Name, studentBody.Email, studentBody.Age)

		if err != nil {
			slog.Error("error updating student user", slog.String("id", id))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]interface{}{
			"message": "Student is updated with these values",
			"student": student,
		})
	}
}

func DeleteById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		slog.Info("deleting a student by id", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}
		deletedId, err := storage.DeleteStudentById(intId)

		if err != nil {
			slog.Error("error deleting student user", slog.String("id", id), slog.Any("deletedId", deletedId))
			response.WriteJson(w, http.StatusInternalServerError, response.GeneralError(err))
			return
		}

		response.WriteJson(w, http.StatusOK, map[string]interface{}{
			"message":    "Student is deleted successfuly",
			"student_id": deletedId,
		})

	}
}
