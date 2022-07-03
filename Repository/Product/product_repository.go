package repository

import (
	"Catalys-Tech-Backend-Test/Model"
	"context"
)

type ProductRepository interface {
	Insert(ctx context.Context, product Model.ProductModel) (Model.ProductModel, error)
	FindById(ctx context.Context, id int32) (Model.ProductModel, error)
	FindAllProductByBrandId(ctx context.Context, id int32) ([]Model.ProductModel, error)
}
