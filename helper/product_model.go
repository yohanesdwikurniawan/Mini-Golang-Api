package helper

import (
	"Catalyst-Tech-Backend-Test/model/domain"
	"Catalyst-Tech-Backend-Test/model/response"
)

func ToProductResponse(product domain.Product) response.ProductResponse {
	return response.ProductResponse{
		Id:      product.Id,
		Name:    product.Name,
		Price:   product.Price,
		BrandId: product.BrandId,
	}
}

func ToProductResponses(products []domain.Product) []response.ProductResponse {
	var productResponses []response.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, ToProductResponse(product))
	}
	return productResponses
}
