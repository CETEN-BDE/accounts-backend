package models

import (
	"accounts/internal/autogen"

	"gorm.io/gorm"
)

type Example struct {
  gorm.Model  // Embeds ID, CreatedAt, UpdatedAt, DeletedAt
  autogen.Example
}