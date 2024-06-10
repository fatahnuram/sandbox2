package main

import (
	"log"
	"net/http"
	"time"
)

func getTime(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(time.Now().String()))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	http.HandleFunc("/time", getTime)

	log.Println("web server started at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("unable to start web server", err)
	}
}
