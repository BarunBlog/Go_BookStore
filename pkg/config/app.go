package config

import "database/sql"

var db *sql.DB

func Connect() {
	d, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	db = d

	// defer the close till after the main function has finished
	// executing
	defer d.Close()
}

func GetDb() *sql.DB {
	return db
}
