package handlers

import (
	"encoding/json"
	"net/http"
	"games_night/server/internal/models"
)

func handleUserCollectionRoutes(w http.ResponseWriter, r *http.Request) {
	mux.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlers.CreateUser(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}\




 func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// Get Firebase UID from context
	token := r.Context().Value("user").(*auth.Token)
	user.FirebaseID = token.UID
	
	if err := models.DB.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(user)
 }