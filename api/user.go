package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/theaaf/todos/app"
	"github.com/theaaf/todos/model"
)

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Id uint `json:"id"`
}

func (a *API) CreateUser(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input UserInput

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	user := &model.User{Email: input.Email}

	if err := ctx.CreateUser(user, input.Password); err != nil {
		return err
	}

	data, err := json.Marshal(&UserResponse{Id: user.ID})
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}
