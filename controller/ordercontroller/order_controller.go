package ordercontroller

import "net/http"

type OrderController interface {
	Insert(writer http.ResponseWriter, request *http.Request)
	FindById(writer http.ResponseWriter, request *http.Request)
}
