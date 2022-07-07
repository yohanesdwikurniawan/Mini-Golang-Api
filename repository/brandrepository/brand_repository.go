package brandrepository

import (
	"Catalyst-Tech-Backend-Test/model/domain"
	"context"
	"database/sql"
)

type BrandRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, brand domain.Brand) (domain.Brand, error)
}
