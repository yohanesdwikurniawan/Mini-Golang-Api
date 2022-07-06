package productservice

import (
	"Catalyst-Tech-Backend-Test/model/request"
	"Catalyst-Tech-Backend-Test/model/response"
	"context"
)

type ProductService interface {
	Insert(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse
	FindById(ctx context.Context, productId int32) response.ProductResponse
	FindAllProductByBrandId(ctx context.Context, brandId int32) []response.ProductResponse
}
