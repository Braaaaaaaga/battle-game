package service

import (
	"battle-game/internal/model"
	"database/sql"
)

func AddCharacters(db *sql.DB, characters []model.Character) error {
	query := `INSERT INTO Characters (user_id, name, health, strength, type, race) VALUES (?, ?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, character := range characters {
		_, err := stmt.Exec(character.UserID, character.Name, character.Health, character.Strength, character.Type, character.Race)
		if err != nil {
			return err
		}
	}
	return nil
}
