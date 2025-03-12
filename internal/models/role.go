package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name string
	IsDefault bool
	Permissions []Permission `gorm:"many2many:role_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	
	Users []*User `gorm:"many2many:user_roles"`
}