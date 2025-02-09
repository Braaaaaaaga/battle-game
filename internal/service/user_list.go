package service

import (
	"battle-game/internal/model"
	"database/sql"
	"encoding/json"
)

func UserList(db *sql.DB) ([]model.User, error) {
	query := `SELECT u.id, u.username, 
        JSON_ARRAYAGG(
            JSON_OBJECT(
                'id', c.id,
                'name', IFNULL(c.name, 'Unknown'),
                'health', IFNULL(c.health, 0),
                'strength', IFNULL(c.strength, 0),
                'type', IFNULL(c.type, 'Unknown'),
                'species', IFNULL(c.species, 'Unknown'),
                'special_attack', IFNULL(c.special_attack, '0'),
                'special_defense', IFNULL(c.special_defense, '0'),
                'defense', IFNULL(c.defense, '0'),
                'weakness', IFNULL(c.weakness, 'None')
            )
        ) AS characters
    FROM users u
    LEFT JOIN characters c ON u.id = c.user_id
    GROUP BY u.id
    ORDER BY u.id ASC;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		var charactersJson string
		if err := rows.Scan(&user.ID, &user.Username, &charactersJson); err != nil {
			return nil, err
		}
		if charactersJson != "" {
			json.Unmarshal([]byte(charactersJson), &user.Characters)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
