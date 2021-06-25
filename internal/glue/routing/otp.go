package routing

import (
	"net/http"

	"Auth/internal/handler/rest"
	"Auth/platform/routers/httprouter"
)

// UserRouting returns the list of routers for domain user
func UserRouting(handler rest.UserHandler) []httprouter.Router {
	return []httprouter.Router{
		{
			Method:  http.MethodPost,
			Path:    "/getotp",
			Handler: handler.Adduser,
		},

		{
			Method:  http.MethodPost,
			Path:    "/validate",
			Handler: handler.Validate,
		},
	}
}
