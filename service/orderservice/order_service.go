package orderservice

import (
	"Catalys-Tech-Backend-Test/model/request"
	"Catalys-Tech-Backend-Test/model/response"
	"context"
)

type OrderService interface {
	Insert(ctx context.Context, request request.OrderCreateRequest) response.OrderResponse
	FindById(ctx context.Context, orderId int32) response.OrderResponse
}
