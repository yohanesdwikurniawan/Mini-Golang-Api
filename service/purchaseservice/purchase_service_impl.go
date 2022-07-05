package purchaseservice

import (
	"Catalys-Tech-Backend-Test/exception"
	"Catalys-Tech-Backend-Test/helper"
	"Catalys-Tech-Backend-Test/model/domain"
	"Catalys-Tech-Backend-Test/model/request"
	"Catalys-Tech-Backend-Test/model/response"
	"Catalys-Tech-Backend-Test/repository/purchaserepository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type PurchaseServiceImpl struct {
	PurchaseRepository purchaserepository.PurchaseRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewPurchaseService(purchaseRepository purchaserepository.PurchaseRepository, DB *sql.DB, validate *validator.Validate) PurchaseService {
	return &PurchaseServiceImpl{
		PurchaseRepository: purchaseRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *PurchaseServiceImpl) Insert(ctx context.Context, request request.PurchaseCreateRequest) response.PurchaseResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	purchase := domain.Purchase{
		ProductId:  request.ProductId,
		BrandId:    request.BrandId,
		ProductQty: request.ProductQty,
		UnitPrice:  request.UnitPrice,
		TotalPrice: request.TotalPrice,
	}

	purchase = service.PurchaseRepository.Insert(ctx, tx, purchase)
	return helper.ToPurchaseResponse(purchase)
}

func (service *PurchaseServiceImpl) FindById(ctx context.Context, purchaseId int32) response.PurchaseResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	purchase, err := service.PurchaseRepository.FindById(ctx, tx, purchaseId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToPurchaseResponse(purchase)
}
