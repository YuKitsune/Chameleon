package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	EmailAddress string
	PasswordHash string

	Aliases []Alias
}
