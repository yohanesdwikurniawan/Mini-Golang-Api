package repository

import (
	"Catalys-Tech-Backend-Test/Model"
	"context"
	"database/sql"
)

type brandRepositoryImpl struct {
	DB *sql.DB
}

func NewBrandRepository(db *sql.DB) BrandRepository {
	return &brandRepositoryImpl{DB: db}
}

func (repository *brandRepositoryImpl) Insert(ctx context.Context, brand Model.BrandModel) (Model.BrandModel, error) {
	script := "INSERT INTO brands(brandName) VALUES (?)"
	result, err := repository.DB.ExecContext(ctx, script, brand.BrandName)
	if err != nil {
		return brand, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return brand, err
	}
	brand.BrandId = int32(id)
	return brand, nil
}
