package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	Connect()
	MakeMigration()
	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", GetTaskById).Methods("GET")
	router.HandleFunc("/tasks/{id}", UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", DeleteTaskById).Methods("DELETE")
	router.HandleFunc("/users", GetAllUsers).Methods("GET")
	router.HandleFunc("/users", AddUser).Methods("POST")
	router.HandleFunc("/users/{id}", GetUserById).Methods("GET")
	router.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/complete_task/{user_id}/{task_id}", UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/fail_task/{user_id}/{task_id}", UpdateUser).Methods("PATCH")

	// start backend server
	fmt.Println("Server starting on port 8000...")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}
