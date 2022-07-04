package brandcontroller

import "net/http"

type BrandController interface {
	Insert(writer http.ResponseWriter, httprequest *http.Request)
}
