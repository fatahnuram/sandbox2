package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().String()))
}

func doMore(w http.ResponseWriter, r *http.Request) {
	fmt.Println("after Handler")
	w.Write([]byte("after Handler"))
}

func main() {
	// HandleFunc can be used when you want/need to control the request life-time
	// i.e. you don't need any chaining/middleware/other processing
	// executes this function and that's it
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	http.HandleFunc("/time", getTime)
	http.HandleFunc("/one", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("nothing more"))
	})

	// Handle can be used when you want/need to pass the executions to the
	// processing chain i.e. middleware before it finally got sent back to client
	http.Handle("/two", func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("inside the Handler first")
			h.ServeHTTP(w, r)
		})
	}(http.HandlerFunc(doMore)))

	log.Println("web server started at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("unable to start web server", err)
	}
}
