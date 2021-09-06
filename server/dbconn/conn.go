package dbconn

import "database/sql"

var Db *sql.DB

func DbConn() (db *sql.DB) {
	var dbase, err = sql.Open("mysql", "root:root@/todolist?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return dbase
}
