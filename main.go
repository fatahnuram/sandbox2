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

func isEven(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if time.Now().Second()%2 == 0 {
			h.ServeHTTP(w, r)
			return
		}
		http.Error(w, "current second is odd, cannot serve the response", http.StatusInternalServerError)
	})
}

func one(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("one\n"))
}
func two(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("two\n"))
}
func three(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("three\n"))
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

	http.HandleFunc("/route/one", one)
	http.HandleFunc("/second/route/here", two)
	http.HandleFunc("/ok/", three)

	// Handle can be used when you want/need to pass the executions to the
	// processing chain i.e. middleware before it finally got sent back to client
	http.Handle("/two", func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("inside the Handler first")
			h.ServeHTTP(w, r)
		})
	}(http.HandlerFunc(doMore)))
	http.Handle("/iseven", isEven(http.HandlerFunc(getTime)))

	log.Println("web server started at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("unable to start web server", err)
	}
}
