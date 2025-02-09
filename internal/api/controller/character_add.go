package controller

import (
	"battle-game/internal/model"
	"battle-game/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddCharacters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var characters []model.Character
	if err := json.NewDecoder(r.Body).Decode(&characters); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i := range characters {
		characters[i].UserID = uint(userID)
	}

	if err := service.AddCharacters(db, characters); err != nil {
		http.Error(w, "Failed to add characters", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Character successfully added to user"))
}
