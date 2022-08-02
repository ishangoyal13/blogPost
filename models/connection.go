package models

import (
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Blog struct {
	ID        uint    `gorm:"primary key:autoIncrement" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Content   *string `json:"content"`
	UserRefer int     `json:"user_id"`
	User      User    `gorm:"foreignKey:UserRefer"`
}

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "host=localhost user=postgres dbname=blogApp password=admin sslmode=disable",
	}))
	if err != nil {
		panic("Error:Failed to connect to database!")
	}

	db.AutoMigrate(&Blog{})
	db.AutoMigrate(&User{})

	DB = db
}
