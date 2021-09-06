package dbops

import (
	"To-Do_App/server/dbconn"
	"To-Do_App/server/dto"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func InsertOneTask(t dto.ToDoList)(c dto.ToDoList){
	task := t.Task
	insForm, err := dbconn.Db.Prepare("INSERT INTO Todolist (task) VALUES(?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(task)
	log.Println("Inserting New Task:",task)

	info,err := dbconn.Db.Query("SELECT * FROM Todolist WHERE id=(SELECT LAST_INSERT_ID())")
	for info.Next() {
		var id int
		var task string
		var status bool
		err = info.Scan(&id, &task, &status)
		if err != nil {
			panic(err.Error())
		}
		c.Id = id
		c.Task = task
		c.Status = status
	}
	return c
}

func DeleteTask(t int){
	delForm, er := dbconn.Db.Prepare("DELETE FROM Todolist WHERE id=?")
	if er != nil {
		panic(er.Error())
	}
	delForm.Exec(t)
	log.Println("ID of Task Deleted:",t)
}

func CompleteOneTask(t int){
	upForm, err := dbconn.Db.Prepare("UPDATE Todolist SET status = true WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	upForm.Exec(t)
	log.Println("ID of Task Completed:",t)
}

func ShowAllTask()(tsks []dto.ToDoList){
	showForm, err := dbconn.Db.Query("SELECT * FROM Todolist ORDER BY id DESC")
	if err != nil {
	panic(err.Error())
	}
	tsk := dto.ToDoList{}
	tsks = []dto.ToDoList{}
	for showForm.Next() {
		var id int
		var task string
		var status bool
		err = showForm.Scan(&id, &task, &status)
		if err != nil {
			panic(err.Error())
		}
		tsk.Id = id
		tsk.Task = task
		tsk.Status = status
		tsks = append(tsks, tsk)
	}
	log.Println("Showing All Tasks:")
	return tsks
}

func ShowTask(t int)( tsk dto.ToDoList){
	showForm, err := dbconn.Db.Query("SELECT * FROM Todolist WHERE id=?",t)
	if err != nil {
		panic(err.Error())
	}
	var id int
	var task string
	var status bool
	for showForm.Next(){
		err = showForm.Scan(&id, &task, &status)
		if err != nil {
			panic(err.Error())
		}
		tsk.Id = id
		tsk.Task = task
		tsk.Status = status
	}
	log.Println("Showing Task with ID",id)
	return tsk
}
