package sandbox2

import (
	"net/http"
	"sandbox2/middleware"
)

type Route struct {
	WithLogger bool
	Handler    http.Handler
}

type Server struct {
	Routes map[string]*Route
}

func New(routes map[string]*Route) *Server {
	return &Server{Routes: routes}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = ShiftPath(r.URL.Path)
	route, ok := s.Routes[head]
	if !ok {
		Respond(w, r, http.StatusNotFound, "root route not found")
		return
	}

	next := route.Handler

	// apply logger middleware
	if route.WithLogger {
		next = middleware.Logger(next)
	}

	next.ServeHTTP(w, r)
}
