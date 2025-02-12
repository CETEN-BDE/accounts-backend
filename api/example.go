package api

import (

	"basic-project/autogen"
	"basic-project/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// (GET /health)
func (s Server) GetHealth(ctx echo.Context) error {

	var example models.Example

	result := s.db.First(&example)
	if result.Error != nil {
		logrus.Errorf("Error getting health: %v", result.Error)
		autogen.GetHealth500JSONResponse{Message: "Error getting health"}.VisitGetHealthResponse(ctx.Response())
		return result.Error
	}

	example.Nb += 1
	result = s.db.Save(&example)
	if result.Error != nil {
		logrus.Errorf("Error updating health: %v", result.Error)
		autogen.GetHealth500JSONResponse{Message: "Error updating health"}.VisitGetHealthResponse(ctx.Response())
		return result.Error
	}

	logrus.Info("/health: Health check")
	autogen.GetHealth200JSONResponse{Status: example.Status, Nb: example.Nb}.VisitGetHealthResponse(ctx.Response())
	return nil
}
