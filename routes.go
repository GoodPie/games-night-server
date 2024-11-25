package main

import (
	"firebase.google.com/go/auth"
	"games_night/server/internal/handlers/rooms"
	"games_night/server/internal/handlers/sessions"
	"games_night/server/internal/handlers/users"
	"games_night/server/internal/middleware"
	"net/http"
)

var mux *http.ServeMux

func setupRoutes(authClient *auth.Client) *http.ServeMux {
	mux = http.NewServeMux()

	// Room routes
	mux.Handle("/api/rooms", middleware.AuthMiddleware(authClient)(http.HandlerFunc(rooms.HandleRoomsCollectionRoutes)))
	mux.Handle("/api/room", middleware.AuthMiddleware(authClient)(http.HandlerFunc(rooms.HandleRoomMemberRoutes)))

	// User routes
	mux.Handle("/api/users", middleware.AuthMiddleware(authClient)(http.HandlerFunc(users.HandleUserCollectionRoutes)))

	// Sessions
	mux.Handle("/api/sessions", middleware.AuthMiddleware(authClient)(http.HandlerFunc(sessions.HandleSessionsCollectionRoutes)))

	return mux
}
