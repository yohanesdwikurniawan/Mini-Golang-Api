package helper

import (
	"Catalyst-Tech-Backend-Test/model/domain"
	"Catalyst-Tech-Backend-Test/model/response"
)

func ToOrderResponse(order domain.Order) response.OrderResponse {
	return response.OrderResponse{
		Id:         order.Id,
		ProductId:  order.ProductId,
		BrandId:    order.BrandId,
		ProductQty: order.ProductQty,
		UnitPrice:  order.UnitPrice,
		TotalPrice: order.TotalPrice,
		CreatedAt:  order.CreatedAt,
	}
}
