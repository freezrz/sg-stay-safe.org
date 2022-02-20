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

/*
export DbUser="admin"
export DbPassword="Tf2pbO26Bc6lqh7AU5jX"
export DbName="sg_stay_safe_db"
export DbHost="sg-stay-safe-db.ciavucegfgwf.ap-southeast-1.rds.amazonaws.com"
export DbPort="3306"
export DbProtocol="tcp"
export DbDriver="mysql"
*/
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
