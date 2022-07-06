package orderrepository

import (
	"Catalyst-Tech-Backend-Test/model/domain"
	"context"
	"database/sql"
)

type OrderRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, order domain.Order) domain.Order
	FindById(ctx context.Context, tx *sql.Tx, orderId int32) (domain.Order, error)
}
