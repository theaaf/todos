package app

import "github.com/theaaf/todos/model"

func (ctx *Context) GetTodoById(id uint) (*model.Todo, error) {
	if ctx.User == nil {
		return nil, ctx.AuthorizationError()
	}

	todo, err := ctx.Database.GetTodoById(id)
	if err != nil {
		return nil, err
	}

	if todo.UserID != ctx.User.ID {
		return nil, ctx.AuthorizationError()
	}

	return todo, nil
}

func (ctx *Context) getTodosByUserId(userId uint) ([]*model.Todo, error) {
	return ctx.Database.GetTodosByUserId(userId)
}

func (ctx *Context) GetUserTodos() ([]*model.Todo, error) {
	if ctx.User == nil {
		return nil, ctx.AuthorizationError()
	}

	return ctx.getTodosByUserId(ctx.User.ID)
}

func (ctx *Context) CreateTodo(todo *model.Todo) error {
	if ctx.User == nil {
		return ctx.AuthorizationError()
	}

	todo.UserID = ctx.User.ID

	if err := ctx.validateTodo(todo); err != nil {
		return err
	}

	return ctx.Database.CreateTodo(todo)
}

const maxTodoNameLength = 100

func (ctx *Context) validateTodo(todo *model.Todo) *ValidationError {
	if len(todo.Name) > maxTodoNameLength {
		return &ValidationError{"name is too long"}
	}

	return nil
}

func (ctx *Context) UpdateTodo(todo *model.Todo) error {
	if ctx.User == nil {
		return ctx.AuthorizationError()
	}

	if todo.UserID != ctx.User.ID {
		return ctx.AuthorizationError()
	}

	if todo.ID == 0 {
		return &ValidationError{"cannot update"}
	}

	if err := ctx.validateTodo(todo); err != nil {
		return nil
	}

	return ctx.Database.UpdateTodo(todo)
}

func (ctx *Context) DeleteTodoById(id uint) error {
	if ctx.User == nil {
		return ctx.AuthorizationError()
	}

	todo, err := ctx.GetTodoById(id)
	if err != nil {
		return err
	}

	if todo.UserID != ctx.User.ID {
		return ctx.AuthorizationError()
	}

	return ctx.Database.DeleteTodoById(id)
}
