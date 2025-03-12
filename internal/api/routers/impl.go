package routers

import (
	"accounts/internal/autogen"

	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

// Ensure the Server type implements the autogen ServerInterface
var _ autogen.ServerInterface = &Server{}

func NewServer(db *gorm.DB) Server {
	return Server{db}
}
