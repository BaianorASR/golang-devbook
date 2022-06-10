package routes

import (
	"api/src/controllers"
	usersController "api/src/controllers/users"
	"net/http"
)

var usersRoute = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		HandlerFunc:  controllers.CreateUser,
		Authrequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		HandlerFunc:  controllers.GetAllUsers,
		Authrequired: false,
	},
	{
		URI:          "/users/search",
		Method:       http.MethodGet,
		HandlerFunc:  usersController.GetUserByNameOrNickname,
		Authrequired: false,
	},
	{
		URI:          "/users/{user_id}",
		Method:       http.MethodGet,
		HandlerFunc:  controllers.GetUserByID,
		Authrequired: false,
	},
	{
		URI:          "/users/{user_id}",
		Method:       http.MethodPut,
		HandlerFunc:  controllers.UpdateUser,
		Authrequired: false,
	},
	{
		URI:          "/users/{user_id}",
		Method:       http.MethodDelete,
		HandlerFunc:  controllers.DeleteUser,
		Authrequired: false,
	},
}
