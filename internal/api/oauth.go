package api

import (
	"accounts/internal/autogen"
	"net/http"

	"github.com/labstack/echo/v4"
)

// (GET /oauth/authorize)
func (s Server) StartAuthorizationFlow(ctx echo.Context, params autogen.StartAuthorizationFlowParams) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}