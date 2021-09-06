package main

import (
	"To-Do_App/server/controller"
	"To-Do_App/server/dbconn"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	dbconn.Db = dbconn.DbConn()
	router:=mux.NewRouter()
	router.HandleFunc("/tasks", controller.GetAllTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/tasks/{id}", controller.GetTask).Methods("GET", "OPTIONS")
	router.HandleFunc("/tasks", controller.CreateTask).Methods("POST", "OPTIONS")
	router.HandleFunc("/tasks/{id}", controller.TaskComplete).Methods("PUT", "OPTIONS")
	router.HandleFunc("/tasks/{id}", controller.DeleteTask).Methods("DELETE", "OPTIONS")
	fmt.Println("Starting server on the port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
	dbconn.Db.Close()
}
