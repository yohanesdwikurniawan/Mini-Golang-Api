package helper

import (
	"Catalys-Tech-Backend-Test/model/domain"
	"Catalys-Tech-Backend-Test/model/response"
)

func ToBrandResponse(brand domain.Brand) response.BrandResponse {
	return response.BrandResponse{
		Id:   brand.Id,
		Name: brand.Name,
	}
}
