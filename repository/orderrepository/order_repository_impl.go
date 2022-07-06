package orderrepository

import (
	"Catalyst-Tech-Backend-Test/helper"
	"Catalyst-Tech-Backend-Test/model/domain"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type OrderRepositoryImpl struct {
}

func NewOrderRepository() OrderRepository {
	return &OrderRepositoryImpl{}
}

func (repository *OrderRepositoryImpl) Insert(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order {
	script := "INSERT INTO orders(productId, brandId, productQty, unitPrice, totalPrice) VALUES (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, script, order.ProductId, order.BrandId, order.ProductQty, order.UnitPrice, order.TotalPrice)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	order.Id = int32(id)
	return order
}

func (repository *OrderRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, orderId int32) (domain.Order, error) {
	script := "SELECT id, productId, brandId, productQty, unitPrice, totalPrice, createdAt FROM orders WHERE id = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, script, orderId)
	helper.PanicIfError(err)
	defer rows.Close()

	order := domain.Order{}
	if rows.Next() {
		err := rows.Scan(&order.Id, &order.ProductId, &order.BrandId, &order.ProductQty, &order.UnitPrice, &order.TotalPrice, &order.CreatedAt)
		helper.PanicIfError(err)
		return order, nil
	} else {
		return order, errors.New("Id " + strconv.Itoa(int(order.Id)) + " Not Found")
	}
}
