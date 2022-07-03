package repository

import (
	"Catalys-Tech-Backend-Test/Model"
	"context"
)

type BrandRepository interface {
	Insert(ctx context.Context, brand Model.BrandModel) (Model.BrandModel, error)
}
