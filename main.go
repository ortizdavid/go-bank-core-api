package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/ortizdavid/go-bank-core-api/config"
	"github.com/ortizdavid/go-bank-core-api/core/controllers"
)

func main() {
	mux := http.NewServeMux()

	dbConn, _ := config.NewDBConnectionFromEnv("DATABASE_URL")
	
	controllers.RegisterRoutes(mux, dbConn.DB)
	// start server
	listenAddr := config.ListenAddr()
	fmt.Printf("Listen to: http://%s\n", listenAddr)
	if err := http.ListenAndServe(listenAddr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}