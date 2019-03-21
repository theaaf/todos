package app

import "github.com/theaaf/todos/model"

func (a *App) GetUserByEmail(email string) (*model.User, error) {
	return a.Database.GetUserByEmail(email)
}
