package brandservice

import (
	"Catalyst-Tech-Backend-Test/model/request"
	"Catalyst-Tech-Backend-Test/model/response"
	"context"
)

type BrandService interface {
	Insert(ctx context.Context, request request.BrandCreateRequest) response.BrandResponse
}
