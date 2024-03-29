package brandrepository

import (
	"Catalyst-Tech-Backend-Test/helper"
	"Catalyst-Tech-Backend-Test/model/domain"
	"context"
	"database/sql"
)

type BrandRepositoryImpl struct {
}

func NewBrandRepository() BrandRepository {
	return &BrandRepositoryImpl{}
}

func (repository *BrandRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, brand domain.Brand) (domain.Brand, error) {
	script := "INSERT INTO brands(name) VALUES (?)"
	result, err := tx.ExecContext(ctx, script, brand.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	brand.Id = int32(id)
	return brand, nil
}
