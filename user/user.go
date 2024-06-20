package user

import (
	"fmt"
	"net/http"
	"sandbox2"
	"strconv"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func New() *sandbox2.Route {
	user := &User{}
	return &sandbox2.Route{
		WithLogger: true,
		Handler:    user,
	}
}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = sandbox2.ShiftPath(r.URL.Path)
	switch head {
	case "":
		u.list(w, r)
	case "echo":
		u.echo(w, r)
	case "demo":
		u.demo(w, r)
	case "detail":
		head, r.URL.Path = sandbox2.ShiftPath(r.URL.Path)
		id, err := strconv.Atoi(head)
		if err != nil {
			sandbox2.Respond(w, r, http.StatusBadRequest, err)
			return
		}
		u.detail(w, r, id)
	default:
		sandbox2.Respond(w, r, http.StatusNotFound, "user path not found")
	}
}

func (u *User) detail(w http.ResponseWriter, r *http.Request, id int) {
	sandbox2.Respond(w, r, http.StatusOK, fmt.Sprintf("detail handler for id %d", id))
}

func (u *User) list(w http.ResponseWriter, r *http.Request) {
	sandbox2.Respond(w, r, http.StatusOK, "list handler")
}

func (u *User) echo(w http.ResponseWriter, r *http.Request) {
	var demo User
	if err := sandbox2.ParseBody(r.Body, &demo); err != nil {
		sandbox2.Respond(w, r, http.StatusBadRequest, err)
		return
	}

	sandbox2.Respond(w, r, http.StatusOK, demo)
}

func (u *User) demo(w http.ResponseWriter, r *http.Request) {
	d := User{Name: "Carlos", Age: 33}
	sandbox2.Respond(w, r, http.StatusOK, d)
}
