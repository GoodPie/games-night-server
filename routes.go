package main

import (
	"firebase.google.com/go/auth"
	"games_night/server/internal/handlers"
	"games_night/server/internal/middleware"
	"net/http"
)

func setupRoutes(authClient *auth.Client) *http.ServeMux {
	mux := http.NewServeMux()

	// Room routes
	mux.Handle("/api/rooms", middleware.AuthMiddleware(authClient)(http.HandlerFunc(handlers.HandleRoomsCollectionRoutes)))
	mux.Handle("/api/room", middleware.AuthMiddleware(authClient)(http.HandlerFunc(handlers.HandleRoomMemberRoutes)))

	// User routes
	mux.Handle("/api/users", middleware.AuthMiddleware(authClient)(http.HandlerFunc(handlers.HandleUserCollectionRoutes)))

	// Sessions
	mux.Handle("/api/sessions", middleware.AuthMiddleware(authClient)(http.HandlerFunc(handlers.HandleSessionsCollectionRoutes)))

	return mux
}
