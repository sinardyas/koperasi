package handler

import (
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func openDbConnection() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root" // Default port if not specified
	}

	dbPasswd := os.Getenv("DB_PASSWORD")
	if dbPasswd == "" {
		dbPasswd = "helloworld" // Default port if not specified
	}

	dbAddr := os.Getenv("DB_HOST")
	if dbAddr == "" {
		dbAddr = "127.0.0.1" // Default port if not specified
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306" // Default port if not specified
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "koperasi" // Default port if not specified
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
