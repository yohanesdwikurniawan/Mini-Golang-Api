package test

import (
	"Catalyst-Tech-Backend-Test/app"
	"Catalyst-Tech-Backend-Test/helper"
	"Catalyst-Tech-Backend-Test/model/domain"
	"Catalyst-Tech-Backend-Test/repository/orderrepository"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOrderInsert(t *testing.T) {
	orderRepository := orderrepository.NewOrderRepository()
	db := app.GetConnection("root", "root", "localhost", "3306", "TestDb")
	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	ctx := context.Background()
	order := domain.Order{
		ProductId: 1,
		BrandId: 1,
		ProductQty: 2,
		UnitPrice: 100000,
		TotalPrice: 200000,
	}

	result, err := orderRepository.Insert(ctx, tx, order)
	helper.PanicIfError(err)

	fmt.Println(result)
}

func TestOrderFindById(t *testing.T) {
	orderRepository := orderrepository.NewOrderRepository()
	db := app.GetConnection("root", "root", "localhost", "3306", "TestDb")
	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	ctx := context.Background()

	result, err := orderRepository.FindById(ctx, tx, 1)
	helper.PanicIfError(err)

	fmt.Println(result)
}