package models

import "gorm.io/gorm"

func InitBlogRepo(db *gorm.DB) IBlog {
	return &blogRepo{
		DB: db,
	}
}

func InitUserRepo(db *gorm.DB) IUser {
	return &userRepo{
		DB: db,
	}
}
