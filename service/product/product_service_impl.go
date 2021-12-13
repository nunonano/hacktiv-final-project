package service

import (
	"errors"

	"github.com/nunonano/hacktiv-final-project/model/entity"
	repository "github.com/nunonano/hacktiv-final-project/repository/product"
)

type ProductServiceImpl struct {
	Repository repository.ProductRepository
}

func InitProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		Repository: productRepository,
	}
}

func (s ProductServiceImpl) GetById(id string) (*entity.Product, error) {
	product := s.Repository.GetById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}

	return product, nil
}

// func (s ProductServiceImpl) GetProduct(id string) (*entity.Product, error) {
// 	product := s.Repository.FindById(id)
// 	if product == nil {
// 		return nil, errors.New("product not found")
// 	}

// 	return product, nil
// }
