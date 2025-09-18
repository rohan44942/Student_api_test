package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rohan44942/student-api/internal/config"
	"github.com/rohan44942/student-api/internal/http/handlers/student"
	"github.com/rohan44942/student-api/internal/storage/sqlite"
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

	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}
	slog.Info("storage initializer", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// router setup
	router := http.NewServeMux()
	// router := mux.NewRouter()

	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/student/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))
	// router.HandleFunc("GET /api/student:id", student.StudentInfo())
	// router.HandleFunc("/api/students", student.New()).Methods("POST")
	// router.HandleFunc("/api/students", student.New()).Methods("POST")
	// router.HandleFunc("/api/student/{id}", student.StudentInfo()).Methods("GET")

	// http server start
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Printf("Server starting on port %s\n", cfg.Addr)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("%v", err)
		}
	}()

	<-done // listing done channale
	// sturctured logs
	slog.Info("shuting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx) // grace fully doing shutdown
	//fmt.Println("Server started on port", cfg.Addr)\
	if err != nil {
		slog.Error("faild to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown successfully")
}
