package main

import (
	"fmt"
	"net/http"
	healthcheck "webapp/services/healthcheck"
)

func main() {
	fmt.Println("Hello, World!")
	startServer()
}

func startServer() {
	fmt.Println("Server started...")
	http.HandleFunc("/healthz", healthcheckHandler)
	http.ListenAndServe(":8080", nil)
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Ensure request body is empty
	if r.ContentLength > 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := healthcheck.Check()
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")

}
