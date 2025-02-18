package models

import (
	"gorm.io/gorm"
)

type Example struct {
  gorm.Model  // Embeds ID, CreatedAt, UpdatedAt, DeletedAt
}