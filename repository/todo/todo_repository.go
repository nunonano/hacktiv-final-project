package repository

import (
	"context"

	"github.com/nunonano/hacktiv-final-project/model/entity"

	"gorm.io/gorm"
)

type TodoRepository interface {
	Delete(ctx context.Context, db *gorm.DB, todoId uint) (uint, error)
	Update(ctx context.Context, db *gorm.DB, todoId uint, todo entity.Todo) entity.Todo
	Create(ctx context.Context, db *gorm.DB, todo entity.Todo) entity.Todo
	GetById(ctx context.Context, db *gorm.DB, todoId uint) (entity.Todo, error)
	GetAll(ctx context.Context, db *gorm.DB) []entity.Todo
}
