package repository

import (
	"context"

	"github.com/nunonano/hacktiv-final-project/model/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(ctx context.Context, db *gorm.DB, order entity.Order) entity.Order
	GetById(ctx context.Context, db *gorm.DB, orderId int) (entity.Order, error)
	GetAll(ctx context.Context, db *gorm.DB) []entity.Order
}
