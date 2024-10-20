package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	GetConnection().Find(&tasks)
	json.NewEncoder(w).Encode(tasks)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)
	TaskId := httpParams["id"]

	var task Task
	GetConnection().Where("id = ?", TaskId).First(&task)
	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	GetConnection().Create(&task)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)
	GetConnection().Updates(&task)
	json.NewEncoder(w).Encode(task)
}

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	httpParams := mux.Vars(r)
	TaskId := httpParams["id"]
	GetConnection().Delete(&User{}, TaskId)
}
