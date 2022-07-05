package brandservice

import (
	"Catalys-Tech-Backend-Test/helper"
	"Catalys-Tech-Backend-Test/model/domain"
	"Catalys-Tech-Backend-Test/model/request"
	"Catalys-Tech-Backend-Test/model/response"
	"Catalys-Tech-Backend-Test/repository/brandrepository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type BrandServiceImpl struct {
	BrandRepository brandrepository.BrandRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewBrandService(brandRepository brandrepository.BrandRepository, DB *sql.DB, validate *validator.Validate) BrandService {
	return &BrandServiceImpl{
		BrandRepository: brandRepository,
		DB:              DB,
		Validate:        validate,
	}
}

func (service *BrandServiceImpl) Insert(ctx context.Context, request request.BrandCreateRequest) response.BrandResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	brand := domain.Brand{
		Name: request.Name,
	}

	brand = service.BrandRepository.Insert(ctx, tx, brand)
	return helper.ToBrandResponse(brand)
}
