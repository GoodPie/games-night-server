package handlers

import (
	"encoding/json"
	"firebase.google.com/go/auth"
	"games_night/server/internal/models"
	"net/http"
)

func HandleSessionsCollectionRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		createSession(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func createSession(w http.ResponseWriter, r *http.Request) {
	var session models.Session
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := r.Context().Value("user").(*auth.Token)

	// Find user by Firebase ID
	var user models.User
	if err := models.DB.Where("firebase_id = ?", token.UID).First(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the creator ID
	session.CreatorId = user.ID

	if err := models.DB.Create(&session).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(session)
}
