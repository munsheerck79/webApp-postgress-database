package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"primarykey;unique;not null;"`
	Name      string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Status    string
	Ownstatus string
	Block     bool
}
