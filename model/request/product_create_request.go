package request

type ProductCreateRequest struct {
	Name    string `validate:"required,min=1,max=100" json:"name"`
	Price   int32  `validate:"required,gte=0" json:"price"`
	BrandId int32  `validate:"required,gte=0" json:"brand_id"`
}
