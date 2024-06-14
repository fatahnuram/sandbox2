package main

import "net/http"

type Route struct {
	Handler http.Handler
}

type App struct {
	User    *Route
	Default *Route
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var next *Route
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	if head == "users" {
		next = app.User
	} else {
		next = app.Default
	}

	next.Handler.ServeHTTP(w, r)
}
