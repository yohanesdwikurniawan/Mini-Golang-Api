package repository

import (
	database "Catalys-Tech-Backend-Test/Database"
	"Catalys-Tech-Backend-Test/Model"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	brandRepository := NewBrandRepository(database.GetConnection("test"))

	ctx := context.Background()
	brand := Model.BrandModel{
		BrandName: "Brand 2",
	}

	result, err := brandRepository.Insert(ctx, brand)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
