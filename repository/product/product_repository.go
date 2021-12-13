package repository

import "github.com/nunonano/hacktiv-final-project/model/entity"

type ProductRepository interface {
	GetById(id string) *entity.Product
}

// Get Data to database
