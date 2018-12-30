package middleware

import (
	"log"
	"net/http"
	"time"
)

func LogInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Time cost: %s in %v", r.URL.Path, time.Since(start))
	})
}