package model

type Todo struct {
	Model

	Name string `json:"name"`
	Done bool   `json:"done"`

	User   User `json:"-"`
	UserID uint `json:"-"`
}
