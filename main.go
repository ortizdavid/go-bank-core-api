package main

import (
	"net/http"
	"github.com/ortizdavid/go-bank-core-api/config"
	"github.com/ortizdavid/go-bank-core-api/controllers"
)

func main() {
	mux := http.NewServeMux()
	controllers.RegisterRoutes(mux)
	http.ListenAndServe(config.ListenAddr(), mux)
}