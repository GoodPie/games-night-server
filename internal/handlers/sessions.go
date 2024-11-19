// internal/handlers/session.go
package handlers

func handleSessionsCollectionRoutes(w http.ResponseWriter, r *http.Request) {
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
	session.AdminID = token.UID

	if err := models.DB.Create(&session).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(session)	
}