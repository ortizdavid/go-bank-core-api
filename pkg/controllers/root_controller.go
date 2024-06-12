package controllers

import (
	"fmt"
	"net/http"
)

func RegisterStaticRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /", index)
	router.HandleFunc("GET /api", index)
	router.HandleFunc("GET /api/download-collections", downloadCollections)
}

func index(w http.ResponseWriter, r *http.Request) {
	html := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Bank Core!</title>
		</head>
		<body>
			<h1>Golang Bank Core API!</h1>
			<p>To test the API in Postman, download the API collections <a target='_blank' href='/api/download-collections'>Here</a>
			</p>
		</body>
		</html>`
	w.Header().Set("Content-Type", "text/html")
    fmt.Fprintln(w, html)
}

func downloadCollections(w http.ResponseWriter, r *http.Request)  {
	fileName := "postman.postman_collection.json"
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, "./_api_collections/"+fileName)
}