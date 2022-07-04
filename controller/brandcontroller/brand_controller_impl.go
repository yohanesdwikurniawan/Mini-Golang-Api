package brandcontroller

import (
	"Catalys-Tech-Backend-Test/helper"
	"Catalys-Tech-Backend-Test/model/request"
	"Catalys-Tech-Backend-Test/model/response"
	"Catalys-Tech-Backend-Test/service/brandservice"
	"net/http"
)

type BrandControllerImpl struct {
	BrandService brandservice.BrandService
}

func NewBrandController(brandService brandservice.BrandService) BrandController {
	return &BrandControllerImpl{
		BrandService: brandService,
	}
}

func (controller *BrandControllerImpl) Insert(writer http.ResponseWriter, httprequest *http.Request) {
	brandCreateRequest := request.BrandCreateRequest{}
	helper.ReadFromRequestBody(httprequest, &brandCreateRequest)

	brandResponse := controller.BrandService.Insert(httprequest.Context(), brandCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   brandResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
