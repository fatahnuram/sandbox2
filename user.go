package main

import "net/http"

type User struct{}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	if head == "" {
		u.home(w, r)
	} else {
		http.Error(w, "user path not found", http.StatusNotFound)
	}
}

func (u *User) home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("user home handler\n"))
}
