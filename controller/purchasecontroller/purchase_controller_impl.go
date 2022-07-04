package purchasecontroller

import (
	"Catalys-Tech-Backend-Test/helper"
	"Catalys-Tech-Backend-Test/model/request"
	"Catalys-Tech-Backend-Test/model/response"
	"Catalys-Tech-Backend-Test/service/purchaseservice"
	"net/http"
	"strconv"
)

type PurchaseControllerImpl struct {
	PurchaseService purchaseservice.PurchaseService
}

func NewPurchaseController(purchaseService purchaseservice.PurchaseService) PurchaseController {
	return &PurchaseControllerImpl{
		PurchaseService: purchaseService,
	}
}

func (controller *PurchaseControllerImpl) Insert(writer http.ResponseWriter, httprequest *http.Request) {
	purchaseCreateRequest := request.PurchaseCreateRequest{}
	helper.ReadFromRequestBody(httprequest, &purchaseCreateRequest)

	purchaseResponse := controller.PurchaseService.Insert(httprequest.Context(), purchaseCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   purchaseResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PurchaseControllerImpl) FindById(writer http.ResponseWriter, httprequest *http.Request) {
	purchaseId := httprequest.URL.Query().Get("purchaseId")
	id, err := strconv.Atoi(purchaseId)
	helper.PanicIfError(err)

	purchaseResponse := controller.PurchaseService.FindById(httprequest.Context(), int32(id))
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   purchaseResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
