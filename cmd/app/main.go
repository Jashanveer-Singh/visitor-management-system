package main

import (
	"log"
	"net/http"

	// _ "github.com/Jashanveer-Singh/visior-management-system" // Change to actual module path
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	// Serve Swagger UI
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	log.Println("Server started at http://localhost:8080/swagger/index.html")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
