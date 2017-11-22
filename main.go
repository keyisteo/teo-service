package main

import (
	"fmt"
	"log"
	"net/http"
	"teo-service/upload"
)

func main() {
	port := 8080
	http.HandleFunc("/upload", upload.Upload)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
