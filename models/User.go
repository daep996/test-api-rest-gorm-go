package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string `gorm:"not null" json:"name"`
	LastName string `gorm:"not null" json:"last_name"`
	Email    string `gorm:"not null;unique_index" json:"email"`
	Tasks    []Task `json:"tasks"`
}
