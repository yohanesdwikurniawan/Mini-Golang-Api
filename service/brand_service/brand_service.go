package brand_service

import (
	"Catalys-Tech-Backend-Test/model/request"
	"Catalys-Tech-Backend-Test/model/response"
	"context"
)

type BrandService interface {
	Insert(ctx context.Context, request request.BrandCreateRequest) response.BrandResponse
}
