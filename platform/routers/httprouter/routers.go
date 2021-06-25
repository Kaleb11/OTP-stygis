package httprouter

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rileyr/middleware"
)

type Routers interface {
	Serve()
}

type Router struct {
	Method      string
	Path        string
	Handler     httprouter.Handle
	MiddleWares []func(handle httprouter.Handle) httprouter.Handle
}
type routing struct {
	host           string
	allowedOrigins string
	domain         string
	port           int
	routers        []Router
}

// Initialize initialize routing and host
func Initialize(host, allowedOrigins string, domain string, routers []Router, port int) Routers {
	return &routing{
		host,
		allowedOrigins,
		domain,
		port,
		routers,
	}
}

func (r *routing) Serve() {
	httpRouter := httprouter.New()
	for _, router := range r.routers {
		if router.MiddleWares == nil {
			httpRouter.Handle(router.Method, router.Path, router.Handler)
		} else {
			s := middleware.NewStack()
			for _, middle := range router.MiddleWares {
				s.Use(middle)
			}
		}
	}
	log.Fatal(http.ListenAndServe(":8080", httpRouter))
}
