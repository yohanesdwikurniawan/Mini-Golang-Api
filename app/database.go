package app

import (
	"Catalys-Tech-Backend-Test/helper"
	"database/sql"
	"embed"
	"time"

	"github.com/pressly/goose/v3"
)

func InitDatabase(name string) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name)
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

func RunMigrations(migrations embed.FS, dbConnection *sql.DB)  {
	
    goose.SetBaseFS(migrations)

    if err := goose.SetDialect("mysql"); err != nil {
        panic(err)
    }

    if err := goose.Up(dbConnection, "migrations"); err != nil {
        panic(err)
    }
}
