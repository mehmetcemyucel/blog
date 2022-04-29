package entity

import "gorm.io/gorm"

type Example struct {
	gorm.Model
	Name string
}
