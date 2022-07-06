package app

import (
	"Catalys-Tech-Backend-Test/controller/brandcontroller"
	"Catalys-Tech-Backend-Test/controller/ordercontroller"
	"Catalys-Tech-Backend-Test/controller/productcontroller"
	"net/http"
)

func AllRouter(brandcontroller brandcontroller.BrandController,
	productcontroller productcontroller.ProductController,
	ordercontroller ordercontroller.OrderController) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/brand", brandcontroller.Insert)
	mux.HandleFunc("/api/product", productcontroller.Insert)
	mux.HandleFunc("/api/product/", productcontroller.FindById)
	mux.HandleFunc("/api/product/brand/", productcontroller.FindAllProductByBrandId)
	mux.HandleFunc("/api/order", ordercontroller.Insert)
	mux.HandleFunc("/api/order/", ordercontroller.FindById)

	return mux
}
