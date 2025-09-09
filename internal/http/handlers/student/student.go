package student

import (
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, bro phla go ka api hit krne ke liye congratulations dude !!!"))
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
