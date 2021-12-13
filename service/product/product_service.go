package service

import "github.com/nunonano/hacktiv-final-project/model/entity"

type ProductService interface {
	GetById(id string) (*entity.Product, error)
}
