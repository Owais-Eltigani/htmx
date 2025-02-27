package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("HTMX server...")

	//
	server := http.NewServeMux()
	// Register static file handler on the custom server mux
	server.Handle("/", http.FileServer(http.Dir("./public")))
	PORT := ":5000"

	server.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// w.Write([]byte("welcome to GO server"))
		json.NewEncoder(w).Encode("success")
	})

	log.Fatal(http.ListenAndServe(PORT, server))
}
