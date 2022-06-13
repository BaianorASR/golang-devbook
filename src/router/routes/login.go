package routes

import (
	loginController "api/src/controllers/login"
	"net/http"
)

var routeLogin = Route{
	URI:          "/login",
	Method:       http.MethodPost,
	HandlerFunc:  loginController.Login,
	Authrequired: false,
}
