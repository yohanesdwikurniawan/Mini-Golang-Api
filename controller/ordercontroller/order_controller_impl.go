package ordercontroller

import (
	"Catalyst-Tech-Backend-Test/helper"
	"Catalyst-Tech-Backend-Test/model/request"
	"Catalyst-Tech-Backend-Test/model/response"
	"Catalyst-Tech-Backend-Test/service/orderservice"
	"net/http"
	"strconv"
)

type OrderControllerImpl struct {
	OrderService orderservice.OrderService
}

func NewOrderController(orderService orderservice.OrderService) OrderController {
	return &OrderControllerImpl{
		OrderService: orderService,
	}
}

func (controller *OrderControllerImpl) Insert(writer http.ResponseWriter, httprequest *http.Request) {
	orderCreateRequest := request.OrderCreateRequest{}
	helper.ReadFromRequestBody(httprequest, &orderCreateRequest)

	orderResponse := controller.OrderService.Insert(httprequest.Context(), orderCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *OrderControllerImpl) FindById(writer http.ResponseWriter, httprequest *http.Request) {
	orderId := httprequest.URL.Query().Get("orderId")
	id, err := strconv.Atoi(orderId)
	helper.PanicIfError(err)

	orderResponse := controller.OrderService.FindById(httprequest.Context(), int32(id))
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
