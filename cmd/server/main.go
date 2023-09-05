package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"golang.a2z/tor_webpage/internal/log"
)

func homePage(logger *zap.Logger) {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		logger.Info("Request from:", zap.String("IP", ip))
		fmt.Fprintf(w, "Hello World")
	}
	http.HandleFunc("/", logFn)
}

func main() {
	logger := log.CreateLogger("info")

	// Define the route handler function for the root URL "/"
	//http.HandleFunc("/", homePage(logger))
	homePage(logger)

	logger.Info("Starting https server")
	logger.Fatal("Error with listen and serve", zap.Error(http.ListenAndServe(":80", nil)))
}
