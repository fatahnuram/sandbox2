package main

import (
	"log"
	"net/http"
	"sandbox2"
	"sandbox2/user"
)

func main() {
	mainRoutes := make(map[string]*sandbox2.Route)

	mainRoutes["users"] = user.New()

	srv := sandbox2.New(mainRoutes)

	log.Println("web server started at localhost:8080")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatal("unable to start web server", err)
	}
}
