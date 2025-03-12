package routers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get all informations about a user.
// (GET /user/{user_id})
func (s Server) GetUser(ctx echo.Context, userId int) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Modify a user
// (PATCH /user/{user_id})
func (s Server) PatchUser(ctx echo.Context, userId int) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Get public information about a user
// (GET /user/{user_id}/profile)
func (s Server) GetUserProfile(ctx echo.Context, userId int) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Remove a role from a user
// (DELETE /user/{user_id}/roles)
func (s Server) DeleteUserRole(ctx echo.Context, userId int) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Get user roles
// (GET /user/{user_id}/roles)
func (s Server) GetUserRoles(ctx echo.Context, userId int) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Give a role to a user
// (POST /user/{user_id}/roles)
func (s Server) PostUserRole(ctx echo.Context, userId int) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Set the roles of a user
// (PUT /user/{user_id}/roles)
func (s Server) PutUserRole(ctx echo.Context, userId int) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Get your own user information
// (GET /userinfo)
func (s Server) GetUserinfo(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Change user information
// (PATCH /userinfo)
func (s Server) PatchUserInfo(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Get a list of all users
// (GET /users)
func (s Server) GetUsers(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Create a user
// (POST /users)
func (s Server) PostUsers(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}
