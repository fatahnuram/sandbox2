package main

import (
	"log"
	"net/http"
	"path"
	"strings"
)

func shiftPath(p string) (string, string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func main() {
	app := &App{
		User:    &Route{Handler: &User{}},
		Default: &Route{Handler: &Default{}},
	}

	log.Println("web server started at localhost:8080")
	if err := http.ListenAndServe(":8080", app); err != nil {
		log.Fatal("unable to start web server", err)
	}
}
