package app

import (
	"Catalyst-Tech-Backend-Test/helper"
	"database/sql"
	"embed"
	"time"

	"github.com/pressly/goose/v3"
)

func InitDatabase(user string, password string, hostname string, port string, dbname string) {
	db, err := sql.Open("mysql", ""+user+":"+password+"@tcp("+hostname+":"+port+")/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		panic(err)
	}
}

func GetConnection(user string, password string, hostname string, port string, dbname string) *sql.DB {
	db, err := sql.Open("mysql", ""+user+":"+password+"@tcp("+hostname+":"+port+")/"+dbname+"?parseTime=true")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func RunMigrations(migrations embed.FS, dbConnection *sql.DB) {

	goose.SetBaseFS(migrations)

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Up(dbConnection, "migrations"); err != nil {
		panic(err)
	}
}
