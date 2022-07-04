package purchaseservice

import (
	"Catalys-Tech-Backend-Test/model/request"
	"Catalys-Tech-Backend-Test/model/response"
	"context"
)

type PurchaseService interface {
	Insert(ctx context.Context, request request.PurchaseCreateRequest) response.PurchaseResponse
	FindById(ctx context.Context, purchaseId int32) response.PurchaseResponse
}
