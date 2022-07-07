package test

import (
	"Catalyst-Tech-Backend-Test/app"
	"Catalyst-Tech-Backend-Test/helper"
	"Catalyst-Tech-Backend-Test/model/domain"
	"Catalyst-Tech-Backend-Test/repository/brandrepository"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestBrandInsert(t *testing.T) {
	brandRepository := brandrepository.NewBrandRepository()
	db := app.GetConnection("root", "root", "localhost", "3306", "TestDb")
	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	
	ctx := context.Background()
	brand := domain.Brand{
		Name : "Brand 2",
	}

	result, err := brandRepository.Insert(ctx, tx, brand)
	helper.PanicIfError(err)
	
	fmt.Println(result)
}