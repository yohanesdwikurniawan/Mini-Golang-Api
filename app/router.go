package app

import (
	"Catalys-Tech-Backend-Test/controller/brandcontroller"
	"Catalys-Tech-Backend-Test/controller/productcontroller"
	"Catalys-Tech-Backend-Test/controller/purchasecontroller"
	"net/http"
)

func AllRouter(brandcontroller brandcontroller.BrandController,
	productcontroller productcontroller.ProductController,
	purchasecontroller purchasecontroller.PurchaseController) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/brand", brandcontroller.Insert)
	mux.HandleFunc("/api/product", productcontroller.Insert)
	mux.HandleFunc("/api/product/", productcontroller.FindById)
	mux.HandleFunc("/api/product/brand/", productcontroller.FindAllProductByBrandId)
	mux.HandleFunc("/api/order", purchasecontroller.Insert)
	mux.HandleFunc("/api/order/", purchasecontroller.FindById)

	return mux
}
