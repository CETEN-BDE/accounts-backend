package routers

import (
	"accounts/internal/autogen"
	"context"
)

// (GET /health)
func (s Server) GetHealth(ctx context.Context, request autogen.GetHealthRequestObject) (autogen.GetHealthResponseObject, error) {
	return autogen.GetHealth200Response{}, nil
}
