package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/theaaf/todos/model"

	"github.com/theaaf/todos/app"
)

func (a *API) GetTodos(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	todos, err := ctx.GetUserTodos()
	if err != nil {
		return err
	}

	data, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

type CreateTodoInput struct {
	Name string `json:"name"`
	Done bool   `json:"done"`
}

type CreateTodoResponse struct {
	Id uint `json:"id"`
}

func (a *API) CreateTodo(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	var input CreateTodoInput

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	todo := &model.Todo{Name: input.Name, Done: input.Done}

	if err := ctx.CreateTodo(todo); err != nil {
		return err
	}

	data, err := json.Marshal(&CreateTodoResponse{Id: todo.ID})
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func (a *API) GetTodoById(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIdFromRequest(r)
	todo, err := ctx.GetTodoById(id)
	if err != nil {
		return err
	}

	data, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

type UpdateTodoInput struct {
	Name *string `json:"name"`
	Done *bool   `json:"done"`
}

func (a *API) UpdateTodoById(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIdFromRequest(r)

	var input UpdateTodoInput

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, &input); err != nil {
		return err
	}

	existingTodo, err := ctx.GetTodoById(id)
	if err != nil || existingTodo == nil {
		return err
	}

	if input.Name != nil {
		existingTodo.Name = *input.Name
	}
	if input.Done != nil {
		existingTodo.Done = *input.Done
	}

	err = ctx.UpdateTodo(existingTodo)
	if err != nil {
		return err
	}

	data, err := json.Marshal(existingTodo)
	if err != nil {
		return err
	}

	_, err = w.Write(data)
	return err
}

func (a *API) DeleteTodoById(ctx *app.Context, w http.ResponseWriter, r *http.Request) error {
	id := getIdFromRequest(r)
	err := ctx.DeleteTodoById(id)

	if err != nil {
		return err
	}

	return &app.UserError{StatusCode: http.StatusOK, Message: "removed"}
}

func getIdFromRequest(r *http.Request) uint {
	vars := mux.Vars(r)
	id := vars["id"]

	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return 0
	}

	return uint(intId)
}
