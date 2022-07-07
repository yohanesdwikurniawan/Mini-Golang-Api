package productrepository

import (
	"Catalyst-Tech-Backend-Test/helper"
	"Catalyst-Tech-Backend-Test/model/domain"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repository *ProductRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, product domain.Product) (domain.Product, error) {
	script := "INSERT INTO products(name, price, brandId) VALUES (?, ?, ?)"
	result, err := tx.ExecContext(ctx, script, product.Name, product.Price, product.BrandId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int32(id)
	return product, nil
}

func (repository *ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int32) (domain.Product, error) {
	script := "SELECT id, name, price, brandId FROM products WHERE id = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, script, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := domain.Product{}
	if rows.Next() {
		// ada
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.BrandId)
		helper.PanicIfError(err)
		return product, nil
	} else {
		// tidak ada
		return product, errors.New("Id " + strconv.Itoa(int(product.Id)) + " Not Found")
	}
}

func (repository *ProductRepositoryImpl) FindAllProductByBrandId(ctx context.Context, tx *sql.Tx, brandId int32) ([]domain.Product, error) {
	script := "SELECT id, name, price, brandId FROM products WHERE brandId = ?"
	rows, err := tx.QueryContext(ctx, script, brandId)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.BrandId)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products, nil
}
