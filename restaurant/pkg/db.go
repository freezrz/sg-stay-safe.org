package pkg

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Database struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
	DbDriver string
	Protocol string
}

var RestaurantDb = Database{
	Host:     "sg-stay-safe.ciavucegfgwf.ap-southeast-1.rds.amazonaws.com",
	Port:     "3306",
	Database: "restaurant_info",
	User:     "admin",
	Password: "whosyourdaddy",
	DbDriver: "mysql",
	Protocol: "tcp",
}

type Restaurant struct {
	Id      int
	Name    string
	Address string
}

func Connect() (db *sql.DB, err error) {
	connString := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",
		RestaurantDb.User,
		RestaurantDb.Password,
		RestaurantDb.Protocol,
		RestaurantDb.Host,
		RestaurantDb.Port,
		RestaurantDb.Database)
	log.Printf("connecting to database: %s", connString)
	db, err = sql.Open(RestaurantDb.DbDriver, connString)
	if err != nil {
		log.Panic(err.Error())
	}
	return db, nil
}

func RetrieveById(id int) *Restaurant {
	result := &Restaurant{}
	db, err := Connect()
	defer func() {
		db.Close()
	}()
	if err != nil {
		log.Panic(err.Error())
	}

	q := fmt.Sprintf(`SELECT id, name, address FROM restaurant WHERE id=%d limit 1;`,
		id)
	log.Println(q)
	r := db.QueryRow(q)
	err = r.Scan(&result.Id, &result.Name, &result.Address)
	if err != nil {
		log.Panic(err.Error())
	}
	return result
}
