package main

import (
	"log"
	"net/http"
)

func main() {
	app := &App{
		User:    &Route{Handler: &User{}},
		Default: &Route{Handler: &Default{}},
	}

	log.Println("web server started at localhost:8080")
	if err := http.ListenAndServe(":8080", withLogger(app)); err != nil {
		log.Fatal("unable to start web server", err)
	}
}
