package db

import (
	"github.com/pkg/errors"

	"github.com/theaaf/todos/model"
)

func (db *Database) GetTodoById(id uint) (*model.Todo, error) {
	var todo model.Todo
	return &todo, errors.Wrap(db.First(&todo, id).Error, "unable to get todo")
}

func (db *Database) GetTodosByUserId(userId uint) ([]*model.Todo, error) {
	var todos []*model.Todo
	return todos, errors.Wrap(db.Find(&todos, model.Todo{UserID: userId}).Error, "unable to get todos")
}

func (db *Database) CreateTodo(todo *model.Todo) error {
	return errors.Wrap(db.Create(todo).Error, "unable to create todo")
}

func (db *Database) UpdateTodo(todo *model.Todo) error {
	return errors.Wrap(db.Save(todo).Error, "unable to update todo")
}

func (db *Database) DeleteTodoById(id uint) error {
	return errors.Wrap(db.Delete(&model.Todo{}, id).Error, "unable to create todo")
}
