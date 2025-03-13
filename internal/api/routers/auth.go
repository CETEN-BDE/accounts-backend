package routers

import (
	"accounts/internal/autogen"
	"context"
)

// Login with username and password
// (POST /auth/password)
func (s Server) ConnectPassword(ctx context.Context, request autogen.ConnectPasswordRequestObject) (autogen.ConnectPasswordResponseObject, error) {
	return autogen.ConnectPassword401Response{}, nil
}
