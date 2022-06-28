package routes

import (
	"devbook-api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		RequireAuth: false,
		Description: "create users",
		Method:      http.MethodPost,
		Handler:     controllers.CreateUser,
	},
	{
		URI:         "/users",
		RequireAuth: false,
		Description: "get all users",
		Method:      http.MethodGet,
		Handler:     controllers.GetAllUsers,
	},
	{
		URI:         "/users/{userId}",
		RequireAuth: false,
		Description: "get user by id",
		Method:      http.MethodGet,
		Handler:     controllers.GetUserById,
	},
	{
		URI:         "/users/{userId}",
		RequireAuth: false,
		Description: "update user by id",
		Method:      http.MethodPut,
		Handler:     controllers.UpdateUser,
	},
	{
		URI:         "/users/{userId}",
		RequireAuth: false,
		Description: "delete user by id",
		Method:      http.MethodDelete,
		Handler:     controllers.DeleteUser,
	},
}
