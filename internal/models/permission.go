package models

import "accounts/internal/autogen"

type Permission struct {
	ID int `gorm:"primarykey"`
	Name autogen.GlobalPermission `gorm:"uniqueIndex"`

	Users []*User `gorm:"many2many:user_permissions;"`
	Roles []*Role `gorm:"many2many:role_permissions;"`
} 
