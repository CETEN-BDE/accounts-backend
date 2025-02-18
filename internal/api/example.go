package api

import (
	"accounts/internal/autogen"

	"github.com/labstack/echo/v4"
)

// (GET /health)
func (s Server) GetHealth(ctx echo.Context) error {
	autogen.GetHealth200Response{}.VisitGetHealthResponse(ctx.Response())
	return nil
}
