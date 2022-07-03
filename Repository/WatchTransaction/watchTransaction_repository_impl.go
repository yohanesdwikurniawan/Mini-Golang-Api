package repository

import (
	"Catalys-Tech-Backend-Test/Model"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type watchTransactionRepositoryImpl struct {
	DB *sql.DB
}

func NewWatchTransactionRepository(db *sql.DB) WatchTransactionRepository {
	return &watchTransactionRepositoryImpl{DB: db}
}

func (repository *watchTransactionRepositoryImpl) Insert(ctx context.Context, watchTransaction Model.WatchTransactionModel) (Model.WatchTransactionModel, error) {
	script := "INSERT INTO watchTransactions(productId, brandId, productQty, unitPrice, totalPrice) VALUES (?, ?, ?, ?, ?)"
	result, err := repository.DB.ExecContext(ctx, script, watchTransaction.ProductId, watchTransaction.BrandId, watchTransaction.ProductQty, watchTransaction.UnitPrice, watchTransaction.TotalPrice)
	if err != nil {
		return watchTransaction, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return watchTransaction, err
	}
	watchTransaction.WatchTransactionId = int32(id)
	return watchTransaction, nil
}

func (repository *watchTransactionRepositoryImpl) FindById(ctx context.Context, id int32) (Model.WatchTransactionModel, error) {
	script := "SELECT productId, brandId, productQty, unitPrice, totalPrice FROM watchTransactions WHERE watchTransactionId = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	watchTransaction := Model.WatchTransactionModel{}
	if err != nil {
		return watchTransaction, err
	}
	defer rows.Close()
	if rows.Next() {
		// ada
		rows.Scan(&watchTransaction.WatchTransactionId, &watchTransaction.ProductId, watchTransaction.BrandId, watchTransaction.ProductQty, watchTransaction.UnitPrice, watchTransaction.TotalPrice, &watchTransaction.CreatedAt)
		return watchTransaction, nil
	} else {
		// tidak ada
		return watchTransaction, errors.New("Id " + strconv.Itoa(int(watchTransaction.WatchTransactionId)) + " Not Found")
	}
}
