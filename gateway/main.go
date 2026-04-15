package main

import (
	"log/slog"
	"net"
	"net/http"
	tokenbucketinmemory "ratelimiter/token-bucket-in-memory"
)

func main() {
	router := http.DefaultServeMux

	rateLimiter := tokenbucketinmemory.NewService()

	router.HandleFunc("/{path...}", func(w http.ResponseWriter, r *http.Request) {
		host, _, _ := net.SplitHostPort(r.RemoteAddr)
		path := r.PathValue("path")

		if ok := rateLimiter.RateLimitRequest(host, r.Method+" /"+path); !ok {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Rate limit reached"))
			return
		}

		http.Redirect(w, r, "http://localhost:8080/"+path, http.StatusFound)
	})

	slog.Info("gateway starting on port 8000")
	if err := http.ListenAndServe(":8000", router); err != nil && err != http.ErrServerClosed {
		slog.Error("gateway error", "error", err)
		return
	}
	slog.Info("gateway closed")

}
