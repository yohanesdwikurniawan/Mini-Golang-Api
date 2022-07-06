package brandcontroller

import (
	"Catalyst-Tech-Backend-Test/helper"
	"Catalyst-Tech-Backend-Test/model/request"
	"Catalyst-Tech-Backend-Test/model/response"
	"Catalyst-Tech-Backend-Test/service/brandservice"
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
