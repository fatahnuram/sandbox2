package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

type StatusResponseWriter struct {
	http.ResponseWriter
	Status       int
	OriginalPath string
}

func (w *StatusResponseWriter) WriteHeader(statuscode int) {
	w.Status = statuscode
	w.ResponseWriter.WriteHeader(statuscode)
}

func (w *StatusResponseWriter) Write(b []byte) (int, error) {
	if w.Status == 0 {
		w.Status = http.StatusOK
	}
	return w.ResponseWriter.Write(b)
}

func shiftPath(p string) (string, string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

func withLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srw := &StatusResponseWriter{
			ResponseWriter: w,
			OriginalPath:   r.URL.Path,
		}

		start := time.Now()
		next.ServeHTTP(srw, r)
		end := time.Since(start)
		fmt.Printf("%s %d %s %v\n", r.Method, srw.Status, srw.OriginalPath, end)
	})
}

func main() {
	app := &App{
		User:    &Route{Handler: &User{}},
		Default: &Route{Handler: &Default{}},
	}

	log.Println("web server started at localhost:8080")
	if err := http.ListenAndServe(":8080", withLogger(app)); err != nil {
		log.Fatal("unable to start web server", err)
	}
}
