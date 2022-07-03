package repository

import (
	"Catalys-Tech-Backend-Test/Model"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type productRepositoryImpl struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepositoryImpl{DB: db}
}

func (repository *productRepositoryImpl) Insert(ctx context.Context, product Model.ProductModel) (Model.ProductModel, error) {
	script := "INSERT INTO products(productName, productPrice, brandId) VALUES (?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, product.ProductName, product.ProductPrice, product.BrandId)
	if err != nil {
		return product, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return product, err
	}
	product.ProductId = int32(id)
	return product, nil
}

func (repository *productRepositoryImpl) FindById(ctx context.Context, id int32) (Model.ProductModel, error) {
	script := "SELECT productId, productName, productPrice, BrandId FROM products WHERE productId = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	product := Model.ProductModel{}
	if err != nil {
		return product, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&product.ProductId, &product.ProductName, &product.ProductPrice, &product.BrandId)
		return product, nil
	} else {
		// tidak ada
		return product, errors.New("Id " + strconv.Itoa(int(product.ProductId)) + " Not Found")
	}
}

func (repository *productRepositoryImpl) FindAllProductByBrandId(ctx context.Context, id int32) ([]Model.ProductModel, error) {
	script := "SELECT productId, productName, productPrice, BrandId FROM products WHERE brandId = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Model.ProductModel
	for rows.Next() {
		product := Model.ProductModel{}
		rows.Scan(&product.ProductId, &product.ProductName, &product.ProductPrice, &product.BrandId)
		products = append(products, product)
	}
	return products, nil
}
