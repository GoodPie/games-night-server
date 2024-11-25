package main

import (
	"games_night/server/internal/firebase"
	"games_night/server/internal/models"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	// Load the environment variables
	saPath := os.Getenv("FIREBASE_SA_PATH") // Path to the service account key JSON file

	// Start Firebase
	firebaseApp, err := firebase.InitFirebase(saPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get Firebase auth client
	auth, err := firebase.GetAuthClient(firebaseApp)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the database
	if err := models.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// Setup routes
	mux := setupRoutes(auth)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
