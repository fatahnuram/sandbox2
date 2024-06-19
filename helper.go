package main

import (
	"encoding/json"
	"io"
	"net/http"
	"path"
	"strings"
)

type CustomError struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

func shiftPath(p string) (string, string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func respond(w http.ResponseWriter, _ *http.Request, status int, data interface{}) error {
	// wrap error as json
	if e, ok := data.(error); ok {
		tmp := CustomError{}
		tmp.Status = "error"
		tmp.Error = e.Error()
		data = tmp
	}

	resp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(resp)

	return nil
}

func parseBody(body io.ReadCloser, result interface{}) error {
	defer body.Close()
	decoder := json.NewDecoder(body)
	return decoder.Decode(result)
}
