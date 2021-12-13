package repository

import (
	"github.com/nunonano/hacktiv-final-project/model/entity"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (r *ProductRepositoryMock) GetById(id string) *entity.Product {
	args := r.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}

	product := args.Get(0).(entity.Product)
	return &product
}
