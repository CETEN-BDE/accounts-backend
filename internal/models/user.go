package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Hidden bool
	GivenName string
	FamilyName string
	Nickname *string
	Picture *string
	Email *string `gorm:"index"`
	EmailVerified bool
	Birthdate sql.NullTime
	PhoneNumber *string
	PhoneNumberVerified bool

	// Username / password authentication
	Username *string `gorm:"index:,options:NULLS NOT DISTINCT"`
	Password *string

	// Google authentication
	GoogleId *string `gorm:"index:,options:NULLS NOT DISTINCT"`

	Roles []*Role `gorm:"many2many:user_roles"`
	Permissions []Permission `gorm:"many2many:user_permissions;"`
}