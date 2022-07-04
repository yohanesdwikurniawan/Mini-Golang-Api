package purchase_repository

import (
	"Catalys-Tech-Backend-Test/helper"
	"Catalys-Tech-Backend-Test/model/domain"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type PurchaseRepositoryImpl struct {
}

func NewPurchaseRepository() PurchaseRepository {
	return &PurchaseRepositoryImpl{}
}

func (repository *PurchaseRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, purchase domain.Purchase) domain.Purchase {
	script := "INSERT INTO purchase(productId, brandId, productQty, unitPrice, totalPrice) VALUES (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, script, purchase.ProductId, purchase.BrandId, purchase.ProductQty, purchase.UnitPrice, purchase.TotalPrice)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	purchase.Id = int32(id)
	return purchase
}

func (repository *PurchaseRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, purchaseId int32) (domain.Purchase, error) {
	script := "SELECT productId, brandId, productQty, unitPrice, totalPrice, createdAt FROM purchase WHERE id = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, script, purchaseId)
	helper.PanicIfError(err)
	defer rows.Close()

	purchase := domain.Purchase{}
	if rows.Next() {
		err := rows.Scan(&purchase.Id, &purchase.ProductId, purchase.BrandId, purchase.ProductQty, purchase.UnitPrice, purchase.TotalPrice, &purchase.CreatedAt)
		helper.PanicIfError(err)
		return purchase, nil
	} else {
		return purchase, errors.New("Id " + strconv.Itoa(int(purchase.Id)) + " Not Found")
	}
}
