package main

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"
)

const (
	CTX_PATH_WIP = "path-is-wip"
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
		wipMsg := r.Context().Value(CTX_PATH_WIP)
		fmt.Printf("%s %d %s %v %v\n", r.Method, srw.Status, srw.OriginalPath, end, wipMsg)
	})
}

func wipContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), CTX_PATH_WIP, "this path is still on progress")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
