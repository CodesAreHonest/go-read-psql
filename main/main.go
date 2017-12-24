package main

import (
	"fmt"
	"database/sql"
	
//	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	

)
	
const (
	DB_USER     = "yinghua"
	DB_PASSWORD = "123"
	DB_NAME     = "fyp1"
)

var db *sql.DB

//====================================================
//function to check error and print error messages
//====================================================
func checkErr(err error, message string) {
	if err != nil {
		panic(message + " err: " + err.Error())
	}
}

//====================================================
// initialize connection with database
//====================================================
func initDB() {

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	sqldb, err := sql.Open("postgres", dbInfo)
	checkErr(err, "Initialize database")
	db = sqldb

}

func main() {
	initDB()
	retrieve_company()
}
