package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID          uint64         `json:"id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string         `json:"name"`
	PhoneNumber int            `json:"phone_number" gorm:"unique"`
	Password    string         `json:"password"`
}

type userRepo struct {
	DB *gorm.DB
}

func (u *userRepo) GetUser(phoneNumber int) (*User, error) {
	return u.GetUserWithTx(u.DB, &User{PhoneNumber: phoneNumber})
}

func (u *userRepo) GetUserWithTx(tx *gorm.DB, where *User) (*User, error) {
	var o User
	err := tx.Model(&User{}).Where(where).Find(&o).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (u *userRepo) CreateUser(user *User) error {
	return u.CreateUserWithTx(u.DB, user)
}

func (u *userRepo) CreateUserWithTx(tx *gorm.DB, user *User) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)

	err = tx.Model(&User{}).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) CheckPassword(providedPassword string, user *User) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
