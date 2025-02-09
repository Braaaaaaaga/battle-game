package model

type Character struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Health         int    `json:"health"`
	Strength       int    `json:"strength"`
	Type           string `json:"type"`
	Species        string `json:"species"`
	SpecialAttack  string `json:"special_attack"`
	SpecialDefense string `json:"special_defense"`
	Defense        string `json:"defense"`
	Weakness       string `json:"weakness"`
	UserID         uint   `json:"-"`
}
