package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var db *gorm.DB

type DB struct {
	*gorm.DB
}

func Init() *gorm.DB {
	// postgres db connection
	err1 := godotenv.Load(".env")
	if err1 != nil {
		log.Fatalf("Error loading .env file")
	}
	user := goDotEnvVariable("USER")
	dbname := goDotEnvVariable("DBNAME")
	password := goDotEnvVariable("PASSWORD")
	host := goDotEnvVariable("HOST")
	port := goDotEnvVariable("PORT")

	var err error
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", user, dbname, password, host, port)
	db, err = gorm.Open(postgres.Open(connStr), &gorm.DB{})

	checkErr(err)

	err = db.Ping()
	checkErr(err)

	return db
}

func goDotEnvVariable(key string) string {
	return os.Getenv(key)
}

// error handling
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
