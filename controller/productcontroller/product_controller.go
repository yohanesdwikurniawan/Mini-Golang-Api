package productcontroller

import "net/http"

type ProductController interface {
	Insert(writer http.ResponseWriter, httprequest *http.Request)
	FindById(writer http.ResponseWriter, httprequest *http.Request)
	FindAllProductByBrandId(writer http.ResponseWriter, httprequest *http.Request)
}
