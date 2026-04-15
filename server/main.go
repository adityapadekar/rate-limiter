package main

import (
	"log/slog"
	"net/http"
)

func newRouter() http.Handler {
	router := http.NewServeMux()

	router.HandleFunc("GET /route1", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request recieved on route 1", "header", r.Header, "remote address", r.RemoteAddr, "request url", r.Pattern)
		w.Write([]byte("Request recieved on route 1\n"))
	})

	router.HandleFunc("GET /route2", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request recieved on route 2", "header", r.Header, "remote address", r.RemoteAddr, "request url", r.Pattern)
		w.Write([]byte("Request recieved on route 2\n"))
	})

	router.HandleFunc("GET /route3", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request recieved on route 3", "header", r.Header, "remote address", r.RemoteAddr, "request url", r.Pattern)
		w.Write([]byte("Request recieved on route 3\n"))
	})

	router.HandleFunc("GET /route4", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request recieved on route 4", "header", r.Header, "remote address", r.RemoteAddr, "request url", r.Pattern)
		w.Write([]byte("Request recieved on route 4\n"))
	})

	return router
}

func main() {

	router := newRouter()

	slog.Info("server starting on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil && err != http.ErrServerClosed {
		slog.Error("server error", "error", err)
		return
	}
	slog.Info("server shutdown")
}
