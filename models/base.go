package models

import "github.com/kim118000/game2023/pkg/db"

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `gorm:"column:created_on" json:"created_on"`
	ModifiedOn int `gorm:"column:modified_on" json:"modified_on"`
	DeletedOn  int `gorm:"column:deleted_on" json:"deleted_on"`
}

func AutoMigrate() {
	db.DB.AutoMigrate(&User{})
}
