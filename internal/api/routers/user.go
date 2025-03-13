package routers

import (
	"accounts/internal/autogen"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get all informations about a user.
// (GET /user/{user_id})
func (s Server) GetUser(ctx context.Context, request autogen.GetUserRequestObject) (autogen.GetUserResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Modify a user
// (PATCH /user/{user_id})
func (s Server) PatchUser(ctx context.Context, request autogen.PatchUserRequestObject) (autogen.PatchUserResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Get public information about a user
// (GET /user/{user_id}/profile)
func (s Server) GetUserProfile(ctx context.Context, request autogen.GetUserProfileRequestObject) (autogen.GetUserProfileResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Remove a role from a user
// (DELETE /user/{user_id}/roles)
func (s Server) DeleteUserRole(ctx context.Context, request autogen.DeleteUserRoleRequestObject) (autogen.DeleteUserRoleResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Get user roles
// (GET /user/{user_id}/roles)
func (s Server) GetUserRoles(ctx context.Context, request autogen.GetUserRolesRequestObject) (autogen.GetUserRolesResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Give a role to a user
// (POST /user/{user_id}/roles)
func (s Server) PostUserRole(ctx context.Context, request autogen.PostUserRoleRequestObject) (autogen.PostUserRoleResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Set the roles of a user
// (PUT /user/{user_id}/roles)
func (s Server) PutUserRole(ctx context.Context, request autogen.PutUserRoleRequestObject) (autogen.PutUserRoleResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Get your own user information
// (GET /userinfo)
func (s Server) GetUserinfo(ctx context.Context, request autogen.GetUserinfoRequestObject) (autogen.GetUserinfoResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Change user information
// (PATCH /userinfo)
func (s Server) PatchUserInfo(ctx context.Context, request autogen.PatchUserInfoRequestObject) (autogen.PatchUserInfoResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Get a list of all users
// (GET /users)
func (s Server) GetUsers(ctx context.Context, request autogen.GetUsersRequestObject) (autogen.GetUsersResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

// Create a user
// (POST /users)
func (s Server) PostUsers(ctx context.Context, request autogen.PostUsersRequestObject) (autogen.PostUsersResponseObject, error) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}
