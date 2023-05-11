package models

import (
	"time"

	"gorm.io/gorm"
)

type Blog struct {
	ID        uint64         `json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	UserID    uint64         `json:"-"`
	User      User           `json:"-"`
	Author    string         `json:"author"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
}

type blogRepo struct {
	DB *gorm.DB
}

func (t *blogRepo) GetBlog(userId uint64) (*[]Blog, error) {
	return t.GetBlogWithTx(t.DB, &Blog{UserID: userId})
}

func (t *blogRepo) GetBlogWithTx(tx *gorm.DB, where *Blog) (*[]Blog, error) {
	var o []Blog
	err := tx.Model(&Blog{}).Where(where).Find(&o).Error
	if err != nil {
		return nil, err
	}
	return &o, nil
}

func (t *blogRepo) CreateBlog(blog *Blog) error {
	return t.CreateBlogWithTx(t.DB, blog)
}

func (t *blogRepo) CreateBlogWithTx(tx *gorm.DB, blog *Blog) error {
	err := tx.Model(&Blog{}).Create(&blog).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *blogRepo) DeleteBlog(id uint64, userId uint64) error {
	return t.DeleteBlogWithTx(t.DB, &Blog{UserID: userId, ID: id})
}

func (t *blogRepo) DeleteBlogWithTx(tx *gorm.DB, where *Blog) error {
	err := tx.Model(&Blog{}).Where(where).Delete(&Blog{}).Error
	if err != nil {
		return err
	}
	return nil
}
