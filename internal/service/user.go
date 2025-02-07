package service

import (
	"battle-game/internal/model"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func RegisterNewUser(db *sql.DB, user *model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err = db.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}
