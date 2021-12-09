package repository

import (
	"context"

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
	err := db.Delete(&todo, todoId).Error
	if err != nil {
		return todoId, err
	}
	return todoId, nil
}

func (repository *TodoRepositoryImpl) Update(ctx context.Context, db *gorm.DB, todoId uint, todo entity.Todo) entity.Todo {
	todo.ID = todoId
	db.Model(&todo).Where("id = ?", todoId).Updates(&todo)
	return todo
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
