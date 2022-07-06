package orderservice

import (
	"Catalys-Tech-Backend-Test/exception"
	"Catalys-Tech-Backend-Test/helper"
	"Catalys-Tech-Backend-Test/model/domain"
	"Catalys-Tech-Backend-Test/model/request"
	"Catalys-Tech-Backend-Test/model/response"
	"Catalys-Tech-Backend-Test/repository/orderrepository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type OrderServiceImpl struct {
	OrderRepository orderrepository.OrderRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewOrderService(orderRepository orderrepository.OrderRepository, DB *sql.DB, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *OrderServiceImpl) Insert(ctx context.Context, request request.OrderCreateRequest) response.OrderResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order := domain.Order{
		ProductId:  request.ProductId,
		BrandId:    request.BrandId,
		ProductQty: request.ProductQty,
		UnitPrice:  request.UnitPrice,
		TotalPrice: request.TotalPrice,
	}

	order = service.OrderRepository.Insert(ctx, tx, order)
	return helper.ToOrderResponse(order)
}

func (service *OrderServiceImpl) FindById(ctx context.Context, orderId int32) response.OrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	order, err := service.OrderRepository.FindById(ctx, tx, orderId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToOrderResponse(order)
}
