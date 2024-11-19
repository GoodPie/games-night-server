package main

import (
	"log"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"games_night/server/internal/models"
	"games_night/server/internal/handlers"
	"games_night/server/internal/firebase"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
 }


func main() {

	// Load the environment variables
	// env := os.Getenv("GO_ENV")
	saPath := os.Getenv("FIREBASE_SA_PATH") // Path to the service account key JSON file

	// Start Firebase
	firebaseApp, err := firebase.InitFirebase(saPath)
	if err != nil {
		log.Fatal(err)
	}
	// Initialize the database
    if err := models.InitDB(); err != nil {
        log.Fatal("Failed to initialize database:", err)
    }
	
	// Start the web server
	mux := http.NewServeMux()

	// Room routes
	mux.Handle("/api/rooms", middleware.AuthMiddleware(auth)(handlers.handleRoomsCollectionRoutes))
	mux.Handle("/api/room", middleware.AuthMiddleware(auth)(handlers.handleRoomMemberRoutes))

	// User routes
	mux.Handle("/api/users", middleware.AuthMiddleware(auth)(handlers.handleUserCollectionRoutes))

	// Sessions
	mux.Handle("/api/sessions", middleware.AuthMiddleware(auth)(handlers.handleSessionsCollectionRoutes))


	// mux.Handle("/ws", middleware.AuthMiddleware(auth)(handlers.HandleWebSocket))

	log.Println("Starting server on :8080")

} 