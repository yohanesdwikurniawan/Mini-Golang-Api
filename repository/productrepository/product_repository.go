package productrepository

import (
	"Catalyst-Tech-Backend-Test/model/domain"
	"context"
	"database/sql"
)

type ProductRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, product domain.Product) (domain.Product, error)
	FindById(ctx context.Context, tx *sql.Tx, productId int32) (domain.Product, error)
	FindAllProductByBrandId(ctx context.Context, tx *sql.Tx, brandId int32) ([]domain.Product, error)
}
