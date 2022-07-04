package purchaserepository

import (
	"Catalys-Tech-Backend-Test/model/domain"
	"context"
	"database/sql"
)

type PurchaseRepository interface {
	Insert(ctx context.Context, tx *sql.Tx, purchase domain.Purchase) domain.Purchase
	FindById(ctx context.Context, tx *sql.Tx, purchaseId int32) (domain.Purchase, error)
}
