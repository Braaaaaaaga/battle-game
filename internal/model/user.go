package model

type User struct {
	ID         uint        `json:"id"`
	Username   string      `json:"username"`
	Password   string      `json:"-"`
	Characters []Character `json:"characters"` // 1 para muitos
}
