package main

import (
	"log"
	"net/http"
)

type mymux struct{}

func (m *mymux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var next http.Handler

	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)

	if head == "one" {
		next = http.HandlerFunc(one)
	} else if head == "two" {
		next = http.HandlerFunc(two)
	} else {
		next = http.HandlerFunc(three)
	}

	next.ServeHTTP(w, r)
}

func shiftPath(path string) (string, string) {
	return "", path
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
	mux := &mymux{}
	log.Println("web server started at localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("unable to start web server", err)
	}
}
