package productservice

import (
	"Catalys-Tech-Backend-Test/exception"
	"Catalys-Tech-Backend-Test/helper"
	"Catalys-Tech-Backend-Test/model/domain"
	"Catalys-Tech-Backend-Test/model/request"
	"Catalys-Tech-Backend-Test/model/response"
	"Catalys-Tech-Backend-Test/repository/productrepository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	ProductRepository productrepository.ProductRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewProductService(productRepository productrepository.ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ProductServiceImpl) Insert(ctx context.Context, request request.ProductCreateRequest) response.ProductResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		Name:    request.Name,
		Price:   request.Price,
		BrandId: request.BrandId,
	}

	product = service.ProductRepository.Insert(ctx, tx, product)
	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindById(ctx context.Context, productId int32) response.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product, err := service.ProductRepository.FindById(ctx, tx, productId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToProductResponse(product)
}

func (service *ProductServiceImpl) FindAllProductByBrandId(ctx context.Context, brandId int32) []response.ProductResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	products := service.ProductRepository.FindAllProductByBrandId(ctx, tx, brandId)

	return helper.ToProductResponses(products)
}
