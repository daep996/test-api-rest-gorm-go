package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/daep996/test-api-rest-gorm-go/db"
	"github.com/daep996/test-api-rest-gorm-go/models"
	"github.com/daep996/test-api-rest-gorm-go/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var PORT = 8080

func main() {
	fmt.Println("Init!")

	errenv := godotenv.Load()
	if errenv != nil {
		log.Fatal("Error loading .env file")
	}

	db.DBConnection()
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HelloHandler)
	// /users routes
	r.HandleFunc("/users", routes.GetUsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users", routes.PostUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods(http.MethodDelete)
	// /task routes
	r.HandleFunc("/tasks", routes.GetTasksHandler).Methods(http.MethodGet)
	r.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods(http.MethodGet)
	r.HandleFunc("/tasks", routes.CreateTaskHandler).Methods(http.MethodPost)
	r.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods(http.MethodDelete)
	// start server
	fmt.Printf("Running on port: %d \n", PORT)
	http.ListenAndServe(":8080", r)
}
