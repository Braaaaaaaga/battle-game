package service

import (
	"battle-game/internal/model"
	"database/sql"
)

func GetRandomCharacter(db *sql.DB, userID uint) (model.Character, error) {
	var character model.Character
	query := `SELECT id, name, health, strength, type, race FROM Characters WHERE user_id=? ORDER BY RAND() LIMIT 1`
	err := db.QueryRow(query, userID).Scan(&character.ID, &character.Name, &character.Health, &character.Strength, &character.Type, &character.Race)
	return character, err
}
