package routers

import (
	"accounts/internal/autogen"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

// (GET /oauth/authorize)
func (s Server) StartAuthorizationFlow(ctx context.Context, request autogen.StartAuthorizationFlowRequestObject) (autogen.StartAuthorizationFlowResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// OAuth token endpoint
	// (POST /oauth/token)
func (s Server) IssueOauthToken(ctx context.Context, request autogen.IssueOauthTokenRequestObject) (autogen.IssueOauthTokenResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}
