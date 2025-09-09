package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohan44942/student-api/internal/config"
	"github.com/rohan44942/student-api/internal/http/handlers/student"
)

func main() {

	// is file me main main char kam krne h
	// 1. config file parse krna h
	// 2. db connection krna h
	// 3. router setup krna h
	// 4. http server start krna h

	// phla kam

	cfg := config.MustLoad() // config file parse ho gyi
	// db connection bad me

	// router setup
	// router := http.NewServeMux()
	router := mux.NewRouter()

	// router.HandleFunc("POST /api/students", student.New())
	// router.HandleFunc("GET /api/student:id", student.StudentInfo())
	// router.HandleFunc("/api/students", student.New()).Methods("POST")
	router.HandleFunc("/api/students", student.New()).Methods("POST")
	router.HandleFunc("/api/student/{id}", student.StudentInfo()).Methods("GET")

	// http server start
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Printf("Server starting on port %s\n", cfg.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("%v", err)
	}
	//fmt.Println("Server started on port", cfg.Addr)

}
