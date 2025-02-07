package model

type Character struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Health   int    `json:"health"`
	Strength int    `json:"strength"`
	Type     string `json:"type"`
	Race     string `json:"race"`
	UserID   uint   `json:"-"`
}
