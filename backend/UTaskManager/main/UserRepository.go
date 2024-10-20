package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	GetConnection().Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)
	UserId := httpParams["id"]
	var user User
	GetConnection().First(&user, "id = ?", UserId)
	json.NewEncoder(w).Encode(user)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	GetConnection().Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	GetConnection().Updates(&user)
	json.NewEncoder(w).Encode(user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)
	UserId := httpParams["id"]
	GetConnection().Delete(&User{}, UserId)
}

// This function are responsible for changing status of an task in database (status of task will be COMPLETED)
// Firstly I update in database the data of a task, then I will will change information for user entity
func CompleteTask(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)
	UserId := httpParams["user_id"]
	TaskId := httpParams["task_id"]

	var task Task
	GetConnection().First(&task, "id = ?", TaskId)
	task.Status = "COMPLETED"
	GetConnection().Updates(&task)

	var user User
	GetConnection().First(&user, "id = ?", UserId)

	userTasks := user.Tasks

	for i := 0; i < len(userTasks); i++ {
		if string(userTasks[i].Id) == TaskId {
			userTasks[i] = task
			break
		}
	}

	user.Tasks = userTasks
	GetConnection().Updates(&user)
}

// This function are responsible for changing status of an task in database (status of task will be FAILED)
// Firstly I update in database the data of a task, then I will will change information for user entity
func RefuseToDoTask(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)
	UserId := httpParams["user_id"]
	TaskId := httpParams["task_id"]

	var task Task
	GetConnection().First(&task, "id = ?", TaskId)
	task.Status = "FAILED"
	GetConnection().Updates(&task)

	var user User
	GetConnection().First(&user, "id = ?", UserId)

	userTasks := user.Tasks

	for i := 0; i < len(userTasks); i++ {
		if string(userTasks[i].Id) == TaskId {
			userTasks[i] = task
			break
		}
	}

	user.Tasks = userTasks
	GetConnection().Updates(&user)
}
