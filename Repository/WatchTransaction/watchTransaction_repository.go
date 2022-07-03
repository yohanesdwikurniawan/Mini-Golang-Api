package repository

import (
	"Catalys-Tech-Backend-Test/Model"
	"context"
)

type WatchTransactionRepository interface {
	Insert(ctx context.Context, watchTransaction Model.WatchTransactionModel) (Model.WatchTransactionModel, error)
	FindById(ctx context.Context, id int32) (Model.WatchTransactionModel, error)
}
