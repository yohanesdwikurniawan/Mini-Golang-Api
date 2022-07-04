package helper

import (
	"Catalys-Tech-Backend-Test/model/domain"
	"Catalys-Tech-Backend-Test/model/response"
)

func ToPurchaseResponse(purchase domain.Purchase) response.PurchaseResponse {
	return response.PurchaseResponse{
		Id:         purchase.Id,
		ProductId:  purchase.ProductId,
		BrandId:    purchase.BrandId,
		ProductQty: purchase.ProductQty,
		UnitPrice:  purchase.UnitPrice,
		TotalPrice: purchase.TotalPrice,
		CreatedAt:  purchase.CreatedAt,
	}
}
