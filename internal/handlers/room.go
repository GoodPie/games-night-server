package handlers

// Route handlers for room routes
// Not sure if this is the best approach but I've been working with Ruby on Rails and this is how I would do it there
func handleRoomsCollectionRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreateRoom(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Routes associated with a specific room
func handleRoomMemberRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetRoom(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


func createRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := models.DB.Create(&room).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(room)
}

func getRoom(w http.ResponseWriter, r *http.Request) {
	var room models.Room
	if err := models.DB.First(&room, mux.Vars(r)["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(room)
}