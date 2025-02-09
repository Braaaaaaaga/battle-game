package service

import (
	"battle-game/internal/model"
	"database/sql"
)

func ReplaceCharacters(db *sql.DB, userID int, characters []model.Character) error {
	_, err := db.Exec("DELETE FROM Characters WHERE user_id = ?", userID)
	if err != nil {
		return err
	}

	query := `INSERT INTO Characters (user_id, name, health, strength, type, species) VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, character := range characters {
		_, err := stmt.Exec(userID, character.Name, character.Health, character.Strength, character.Type, character.Species)
		if err != nil {
			return err
		}
	}
	return nil
}
