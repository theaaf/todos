package db

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/theaaf/todos/model"
)

func (db *Database) GetTodoById(id uint) (*model.Todo, error) {
	var todo model.Todo

	if err := db.First(&todo, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "unable to get todo")
	}

	return &todo, nil
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
	return errors.Wrap(db.Delete(&model.Todo{}, id).Error, "unable to delete todo")
}
