package request

type OrderCreateRequest struct {
	ProductId  int32 `validate:"required,gte=0" json:"product_id"`
	BrandId    int32 `validate:"required,gte=0" json:"brand_id"`
	ProductQty int32 `validate:"required,gte=0" json:"product_qty"`
	UnitPrice  int32 `validate:"required,gte=0" json:"unit_price"`
	TotalPrice int32 `validate:"required,gte=0" json:"total_price"`
}
