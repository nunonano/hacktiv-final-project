package repository

import (
	"context"
	"errors"

	"github.com/nunonano/hacktiv-final-project/model/entity"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
}

func InitCategoryRepository() TodoRepository {
	return &TodoRepositoryImpl{}
}

func (repository *TodoRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, todoId uint) (uint, error) {
	var todo entity.Todo
	if db.Delete(&todo, todoId).RowsAffected == 0 {
		return todoId, errors.New("Delete Failed")
	}
	return todoId, nil
}

func (repository *TodoRepositoryImpl) Update(ctx context.Context, db *gorm.DB, todoId uint, todo entity.Todo) (entity.Todo, error) {
	todo.ID = todoId
	if db.Model(&todo).Where("id = ?", todoId).Updates(&todo).RowsAffected == 0 {
		return todo, errors.New("Update Failed")
	}
	return todo, nil
}

func (repository *TodoRepositoryImpl) Create(ctx context.Context, db *gorm.DB, todo entity.Todo) entity.Todo {
	db.Save(&todo)
	return todo
}

func (repository *TodoRepositoryImpl) GetById(ctx context.Context, db *gorm.DB, todoId uint) (entity.Todo, error) {
	var todo entity.Todo
	err := db.First(&todo, todoId).Error
	if err != nil {
		return todo, err
	}
	return todo, nil
}

func (repository *TodoRepositoryImpl) GetAll(ctx context.Context, db *gorm.DB) []entity.Todo {
	var todos []entity.Todo
	db.Find(&todos)
	return todos
}
