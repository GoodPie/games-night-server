package handlers

import (
	"encoding/json"
	"firebase.google.com/go/auth"
	"games_night/server/internal/models"
	"net/http"
)

func HandleUserCollectionRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get Firebase UID from context
	token := r.Context().Value("user").(*auth.Token)
	user.FirebaseId = token.UID

	if err := models.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}
