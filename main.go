package main

import (
	"fmt"
	"net/http"

	"github.com/jerome-wilson/GO-REST-API/handlers"
)

func main() {
	fmt.Println("Server running at http://localhost:3000")
	http.HandleFunc("/books/", handlers.HandleBooks)
	http.ListenAndServe(":3000", nil)
}
