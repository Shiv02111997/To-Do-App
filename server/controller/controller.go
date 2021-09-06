package controller

import (
	"To-Do_App/server/dbops"
	"To-Do_App/server/dto"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)
// Sending JSON Response

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(result)
}


func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task dto.ToDoList
	_ = json.NewDecoder(r.Body).Decode(&task)
	t:= dbops.InsertOneTask(task)
	respondWithJSON(w,http.StatusOK,t)
}


func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	f,_:=strconv.Atoi(id)
	var task dto.ToDoList
	task= dbops.ShowTask(f)
	dbops.DeleteTask(f)
	respondWithJSON(w,http.StatusOK,task)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	t,_:= strconv.Atoi(id)
	dbops.CompleteOneTask(t)
	task := dto.ToDoList{}
	task= dbops.ShowTask(t)
	respondWithJSON(w,http.StatusOK,task)
}

func GetAllTask(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w,http.StatusOK, dbops.ShowAllTask())
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	t,_:= strconv.Atoi(id)
	respondWithJSON(w,http.StatusOK, dbops.ShowTask(t))
}
