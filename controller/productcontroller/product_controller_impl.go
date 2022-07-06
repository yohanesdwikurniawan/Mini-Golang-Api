package productcontroller

import (
	"Catalyst-Tech-Backend-Test/helper"
	"Catalyst-Tech-Backend-Test/model/request"
	"Catalyst-Tech-Backend-Test/model/response"
	"Catalyst-Tech-Backend-Test/service/productservice"
	"net/http"
	"strconv"
)

type ProductControllerImpl struct {
	ProductService productservice.ProductService
}

func NewProductController(productService productservice.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: productService,
	}
}

func (controller *ProductControllerImpl) Insert(writer http.ResponseWriter, httprequest *http.Request) {
	productCreateRequest := request.ProductCreateRequest{}
	helper.ReadFromRequestBody(httprequest, &productCreateRequest)

	productResponse := controller.ProductService.Insert(httprequest.Context(), productCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, httprequest *http.Request) {
	productId := httprequest.URL.Query().Get("productId")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.FindById(httprequest.Context(), int32(id))
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindAllProductByBrandId(writer http.ResponseWriter, httprequest *http.Request) {
	brandId := httprequest.URL.Query().Get("brandId")
	id, err := strconv.Atoi(brandId)
	helper.PanicIfError(err)

	productResponse := controller.ProductService.FindAllProductByBrandId(httprequest.Context(), int32(id))
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   productResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
