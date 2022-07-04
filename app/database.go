package app

import (
	"Catalys-Tech-Backend-Test/helper"
	"database/sql"
	"time"
)

func InitDatabase(name string) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		panic(err)
	}
}

func GetConnection(name string) *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/"+name+"?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
