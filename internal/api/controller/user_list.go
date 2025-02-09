package controller

import (
	"battle-game/internal/service"
	"encoding/json"
	"net/http"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.UserList(db)
	if err != nil {
		http.Error(w, "Failed to get users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
