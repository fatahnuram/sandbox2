package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	CTX_PATH_WIP         = "path-is-wip"
	CTX_LOGGER_START     = "engine-logger-start"
	CTX_LOGGER_ORIG_PATH = "engine-logger-path"
	CTX_LOGGER_STATUS    = "engine-logger-status"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ctx := context.WithValue(r.Context(), CTX_LOGGER_START, start)
		ctx = context.WithValue(ctx, CTX_LOGGER_ORIG_PATH, r.URL.Path)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func LogRequest(r *http.Request, status int) {
	ctx := r.Context()

	path, _ := ctx.Value(CTX_LOGGER_ORIG_PATH).(string)
	start, _ := ctx.Value(CTX_LOGGER_START).(time.Time)

	fmt.Printf("%s %d %s %v\n", r.Method, status, path, time.Since(start))
}

func wipContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), CTX_PATH_WIP, "this path is still on progress")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
