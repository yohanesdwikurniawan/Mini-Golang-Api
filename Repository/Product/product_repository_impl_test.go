package repository

import (
	database "Catalys-Tech-Backend-Test/Database"
	"Catalys-Tech-Backend-Test/Model"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestProductInsert(t *testing.T) {
	productRepository := NewProductRepository(database.GetConnection("test"))

	ctx := context.Background()
	product := Model.ProductModel{
		ProductName:  "Product 2.3",
		ProductPrice: 250000,
		BrandId:      2,
	}

	result, err := productRepository.Insert(ctx, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	productRepository := NewProductRepository(database.GetConnection("test"))

	product, err := productRepository.FindById(context.Background(), 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
}

func TestFindAll(t *testing.T) {
	productRepository := NewProductRepository(database.GetConnection("test"))

	products, err := productRepository.FindAllProductByBrandId(context.Background(), 1)
	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Println(product)
	}
}
