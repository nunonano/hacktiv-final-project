package service_test

import (
	"github.com/nunonano/hacktiv-final-project/model/entity"
	repository "github.com/nunonano/hacktiv-final-project/repository/product"
	service "github.com/nunonano/hacktiv-final-project/service/product"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"testing"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = service.InitProductService(productRepository)

func TestProductService_GetProductNotFound(t *testing.T) {
	// id 1 --> data == nil

	//Mocking data
	productRepository.Mock.On("GetById", "1").Return(nil)

	product, err := productService.GetById("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)

	assert.Equal(t, "product not found", err.Error(), "error should product not found")
}

func TestProductService_GetProduct(t *testing.T) {
	// id 2 --> data == product
	productMock := entity.Product{
		Id:   "2",
		Name: "Product 2",
	}

	// Mocking data
	productRepository.Mock.On("GetById", "2").Return(productMock)

	product, err := productService.GetById("2")

	assert.Nil(t, err)

	assert.NotNil(t, product)

	assert.Equal(t, &productMock, product, "result must be equals")
}
