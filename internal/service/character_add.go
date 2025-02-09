package service

import (
	"battle-game/internal/model"
	"database/sql"
)

func AddCharacters(db *sql.DB, characters []model.Character) error {
	query := `INSERT INTO Characters (user_id, name, health, strength, type, species, special_attack, special_defense, defense, weakness) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, character := range characters {
		_, err := stmt.Exec(character.UserID, character.Name, character.Health, character.Strength, character.Type, character.Species, character.SpecialAttack, character.SpecialDefense, character.Defense, character.Weakness)
		if err != nil {
			return err
		}
	}
	return nil
}
