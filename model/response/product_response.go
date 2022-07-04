package response

type ProductResponse struct {
	Id      int32  `json:"id"`
	Name    string `json:"name"`
	Price   int32  `json:"price"`
	BrandId int32  `json:"brand_id"`
}
