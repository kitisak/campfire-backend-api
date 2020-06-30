package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheckHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}
