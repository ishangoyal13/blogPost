package models

import "gorm.io/gorm"

type IBlog interface {
	GetBlog(userId uint64) (*[]Blog, error)
	GetBlogWithTx(tx *gorm.DB, where *Blog) (*[]Blog, error)
	CreateBlog(blog *Blog) error
	CreateBlogWithTx(tx *gorm.DB, blog *Blog) error
	DeleteBlog(id uint64, userId uint64) error
	DeleteBlogWithTx(tx *gorm.DB, where *Blog) error
}

type IUser interface {
	GetUser(phoneNumber int) (*User, error)
	GetUserWithTx(tx *gorm.DB, where *User) (*User, error)
	CreateUser(user *User) error
	CreateUserWithTx(tx *gorm.DB, user *User) error
	CheckPassword(providedPassword string, user *User) error
}
