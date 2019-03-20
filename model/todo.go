package model

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model

	Name string
	Done bool

	User   User
	UserID uint
}
