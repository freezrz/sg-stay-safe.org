package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Restaurant struct {
	Id      int
	Name    string
	Address string
}

func Connect() (db *sql.DB, err error) {
	connString := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		os.Getenv("DbUser"),
		os.Getenv("DbPassword"),
		os.Getenv("DbProtocol"),
		os.Getenv("DbHost"),
		os.Getenv("DbPort"),
		os.Getenv("DbName"))
	log.Printf("connecting to database: %s", connString)
	db, err = sql.Open(os.Getenv("DbDriver"), connString)
	if err != nil {
		log.Panic(err.Error())
	}
	return db, nil
}
