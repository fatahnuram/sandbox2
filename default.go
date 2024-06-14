package main

import "net/http"

type Default struct{}

func (d *Default) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome!\n"))
}
