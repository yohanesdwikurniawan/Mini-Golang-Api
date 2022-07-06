package helper

import (
	"Catalyst-Tech-Backend-Test/model/domain"
	"Catalyst-Tech-Backend-Test/model/response"
)

func ToBrandResponse(brand domain.Brand) response.BrandResponse {
	return response.BrandResponse{
		Id:   brand.Id,
		Name: brand.Name,
	}
}
