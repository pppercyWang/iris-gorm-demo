package models

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gorm:"type:varchar(20);not null;"`
	Count        string    `gorm:"type:varchar(10);not null;"`
	Author        string    `gorm:"type:varchar(20);not null;"`
	Type     string `gorm:"type:varchar(20);not null;"`
}


