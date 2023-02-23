package main

import (
	"github.com/gorilla/mux"
	"github.com/khalil-farashiani/OrderPalace/sender/internal/adapters/broker/redis"
	v1 "github.com/khalil-farashiani/OrderPalace/sender/internal/delivery/http/v1"
	"net/http"
)

func setHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// set the necessary headers
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// call the next middleware/handler in the chain
		next.ServeHTTP(w, r)
	})
}

func getRoutes(redis redis.RedisQueue) *mux.Router {
	r := mux.NewRouter()
	r.Use(setHeadersMiddleware)
	r.HandleFunc("/api/order/", v1.SendOrder(redis))

	return r
}
