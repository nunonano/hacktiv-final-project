package service

import (
	"context"

	api "github.com/nunonano/hacktiv-final-project/model/api/order"
	"github.com/nunonano/hacktiv-final-project/model/entity"
	repository "github.com/nunonano/hacktiv-final-project/repository/order"

	"gorm.io/gorm"
)

type OrderServiceImpl struct {
	OrderRepository repository.OrderRepository
	DB              *gorm.DB
}

func InitOrderService(orderRepository repository.OrderRepository, DB *gorm.DB) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB:              DB,
	}
}

// @Tags order
// @Summary Create Order
// @Description Create Order Description
// @Accept json
// @Produce json
// @Success 200 {array} entity.Order
// @Param id body entity.Order true "Input Data"
// @Router /order [post]
func (service *OrderServiceImpl) Create(ctx context.Context, request api.OrderCreateRequest) api.OrderResponse {

	order := entity.Order{
		Name: request.Name,
	}

	order = service.OrderRepository.Create(ctx, service.DB, order)

	return api.OrderResponse{
		Name: order.Name,
		Id:   order.Id,
	}
}

// @Tags order
// @Summary Get All Order
// @Description This is information about all orders
// @Accept json
// @Produce json
// @Success 200 {array} entity.Order
// @Router /order [get]
func (service *OrderServiceImpl) GetAll(ctx context.Context) []api.OrderResponse {

	var orderResponses []api.OrderResponse

	orders := service.OrderRepository.GetAll(ctx, service.DB)

	for _, order := range orders {
		orderResponses = append(orderResponses, api.OrderResponse{Name: order.Name,
			Id: order.Id})
	}

	return orderResponses
}

// @Tags order
// @Summary Get Order by ID
// @Description This is information about order
// @Accept json
// @Produce json
// @Success 200 {array} entity.Order
// @Param id path int true "Order ID"
// @Router /order/{id} [get]
func (service *OrderServiceImpl) GetById(ctx context.Context, orderId int) api.OrderResponse {
	order, err := service.OrderRepository.GetById(ctx, service.DB, orderId)
	if err != nil {
		return api.OrderResponse{}
	} else {
		return api.OrderResponse{
			Name: order.Name,
			Id:   order.Id,
		}
	}
}
