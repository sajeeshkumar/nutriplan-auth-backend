package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var users map[string]string

func loadUsers() {
	users = make(map[string]string)

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	envUsers := os.Getenv("USERS")

	if envUsers == "" {
		log.Fatal("USERS environment variable is not set")
	}

	for _, pair := range strings.Split(envUsers, ",") {
		parts := strings.Split(pair, ":")
		if len(parts) == 2 {
			users[parts[0]] = parts[1]
		}
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		fmt.Println("Error decoding JSON:", err)
		return
	}

	if password, exists := users[creds.Username]; exists {
		if password == creds.Password {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Login successful"})
			return
		} else {
			fmt.Println("Password mismatch")
		}
	} else {
		fmt.Println("User not found:", creds.Username)
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

func main() {
	loadUsers()
	router := mux.NewRouter()

	// Your routes
	router.HandleFunc("/login", authenticate).Methods(http.MethodPost)

	// CORS setup
	c := cors.New(cors.Options{
        AllowedOrigins: []string{
            "http://localhost:3000",      // Local development
            "https://nutriplan-frontend.onrender.com", // Render deployment
        },
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization", "Referrer-Policy"},
        AllowCredentials: true, // If you use cookies or Authorization headers
    })

	// Wrap your router with CORS middleware
	handler := c.Handler(router)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
