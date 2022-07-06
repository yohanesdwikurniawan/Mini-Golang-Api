package domain

import "time"

type Order struct {
	Id         int32
	ProductId  int32
	BrandId    int32
	ProductQty int32
	UnitPrice  int32
	TotalPrice int32
	CreatedAt  time.Time
}
