package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:name`
	Age  int    `json:age`
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)
	if head == "" {
		u.home(w, r)
	} else if head == "echo" {
		u.echo(w, r)
	} else if head == "demo" {
		u.demo(w, r)
	} else {
		http.Error(w, "user path not found", http.StatusNotFound)
	}
}

func (u *User) home(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("user home handler\n"))
}

func (u *User) echo(w http.ResponseWriter, r *http.Request) {
	var demo User
	if err := json.NewDecoder(r.Body).Decode(&demo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	w.Write([]byte(fmt.Sprintf("%s is %d years old\n", demo.Name, demo.Age)))
}

func (u *User) demo(w http.ResponseWriter, _ *http.Request) {
	d := User{Name: "Carlos", Age: 33}
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
