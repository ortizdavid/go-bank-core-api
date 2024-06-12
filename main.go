package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/ortizdavid/go-bank-core-api/config"
	"github.com/ortizdavid/go-bank-core-api/pkg/controllers"
)

func main() {
	mux := http.NewServeMux()

	// establish database connections
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer config.DisconnectDB(db)
	
	// setup routes
	controllers.RegisterStaticRoutes(mux)
	controllers.RegisterRoutes(mux, db)

	// start server
	listenAddr := config.ListenAddr()
	fmt.Printf("Listen to: http://%s\n", listenAddr)
	if err := http.ListenAndServe(listenAddr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}