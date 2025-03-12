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
	Nickname sql.NullString
	Picture sql.NullString
	Email sql.NullString `gorm:"index"`
	EmailVerified bool
	Birthdate sql.NullTime
	PhoneNumber sql.NullString
	PhoneNumberVerified bool

	// Username / password authentication
	Username sql.NullString `gorm:"index:,options:NULLS NOT DISTINCT"`
	Password sql.NullString

	// Google authentication
	GoogleId sql.NullString `gorm:"index:,options:NULLS NOT DISTINCT"`

	Roles []*Role `gorm:"many2many:user_roles"`
	Permissions []Permission `gorm:"many2many:user_permissions;"`
}