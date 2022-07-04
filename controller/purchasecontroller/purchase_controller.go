package purchasecontroller

import "net/http"

type PurchaseController interface {
	Insert(writer http.ResponseWriter, request *http.Request)
	FindById(writer http.ResponseWriter, request *http.Request)
}
