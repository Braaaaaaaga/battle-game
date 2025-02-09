package controller

import (
	"battle-game/internal/service"
	"encoding/json"
	"net/http"
)

func BattleHandler(w http.ResponseWriter, r *http.Request) {
	type BattleRequest struct {
		UserID1 uint `json:"user_id_1"`
		UserID2 uint `json:"user_id_2"`
	}

	var req BattleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	char1, err := service.GetRandomCharacter(db, req.UserID1)
	if err != nil {
		http.Error(w, "Failed to retrieve character for user 1", http.StatusInternalServerError)
		return
	}

	char2, err := service.GetRandomCharacter(db, req.UserID2)
	if err != nil {
		http.Error(w, "Failed to retrieve character for user 2", http.StatusInternalServerError)
		return
	}

	result := "Draw"
	if char1.Strength > char2.Strength {
		result = "User 1 wins"
	} else if char2.Strength > char1.Strength {
		result = "User 2 wins"
	}

	response := map[string]string{
		"result":     result,
		"character1": char1.Name,
		"character2": char2.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
