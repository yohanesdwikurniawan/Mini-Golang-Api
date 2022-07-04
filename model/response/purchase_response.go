package response

import "time"

type PurchaseResponse struct {
	Id         int32     `json:"id"`
	Name       string    `json:"name"`
	ProductId  int32     `json:"product_id"`
	BrandId    int32     `json:"brand_id"`
	ProductQty int32     `json:"product_qty"`
	UnitPrice  int32     `json:"unit_price"`
	TotalPrice int32     `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
}
