package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"username"`
	Email string `json:"email"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     struct {
		Lat string `json:"lat"`
		Lng string `json:"lng"`
	} `json:"geo"`
}

type Users struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
}

var users = []User{
	{ID: "1", Name: "John Doe", Email: "john@example.com"},
	{ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
	{ID: "3", Name: "Bob Wilson", Email: "bob@example.com"},
	{ID: "4", Name: "Alice Brown", Email: "alice@example.com"},
	{ID: "5", Name: "Charlie Davis", Email: "charlie@example.com"},
	{ID: "6", Name: "Eva Miller", Email: "eva@example.com"},
	{ID: "7", Name: "Frank Johnson", Email: "frank@example.com"},
	{ID: "8", Name: "Grace Lee", Email: "grace@example.com"},
	{ID: "9", Name: "Henry Taylor", Email: "henry@example.com"},
	{ID: "10", Name: "Ivy Chen", Email: "ivy@example.com"},
}

func main() {
	fmt.Println("HTMX server...")

	//
	server := http.NewServeMux()
	// Register static file handler on the custom server mux
	server.Handle("/", http.FileServer(http.Dir("./public")))
	PORT := ":5000"

	server.HandleFunc("GET /welcome", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// w.Write([]byte("welcome to GO server"))
		json.NewEncoder(w).Encode("success")
	})

	server.HandleFunc("GET /users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(users)
	})

	server.HandleFunc("GET /other-users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		var externalUsers []Users
		if err := json.NewDecoder(resp.Body).Decode(&externalUsers); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(externalUsers)
		fmt.Println(externalUsers)
	})

	server.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		id := r.URL.Query().Get("id")
		fmt.Println("id: ", id)

		var wantedUser = User{}
		for _, user := range users {
			if user.ID == id {

				wantedUser = user
				json.NewEncoder(w).Encode(wantedUser)
				return
			}
		}
		http.Error(w, "not found", http.StatusNotFound)
		w.Write([]byte("user not found"))

	})

	log.Fatal(http.ListenAndServe(PORT, server))
}
