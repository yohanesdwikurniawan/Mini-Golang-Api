package test

import (
	"Catalyst-Tech-Backend-Test/app"
	"Catalyst-Tech-Backend-Test/helper"
	"Catalyst-Tech-Backend-Test/model/domain"
	"Catalyst-Tech-Backend-Test/repository/productrepository"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestProductInsert(t *testing.T) {
	productRepository := productrepository.NewProductRepository()
	db := app.GetConnection("root", "root", "localhost", "3306", "TestDb")
	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	ctx := context.Background()
	product := domain.Product{
		Name: "Test Product 1",
		Price: 100000,
		BrandId: 1,
	}

	result, err := productRepository.Insert(ctx, tx, product)
	helper.PanicIfError(err)

	fmt.Println(result)
}

func TestProductFindById(t *testing.T) {
	productRepository := productrepository.NewProductRepository()
	db := app.GetConnection("root", "root", "localhost", "3306", "TestDb")
	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	ctx := context.Background()

	result, err := productRepository.FindById(ctx, tx, 1)
	helper.PanicIfError(err)

	fmt.Println(result)
}

func TestProductFindProductsByBrandId(t *testing.T) {
	productRepository := productrepository.NewProductRepository()
	db := app.GetConnection("root", "root", "localhost", "3306", "TestDb")
	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	ctx := context.Background()

	result,err := productRepository.FindAllProductByBrandId(ctx, tx, 1)
	helper.PanicIfError(err)
	
	fmt.Println(result)
}