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

	mux.HandleFunc("/api/brands", brandcontroller.Insert)
	mux.HandleFunc("/api/products", productcontroller.Insert)
	mux.HandleFunc("/api/products/", productcontroller.FindById)
	mux.HandleFunc("/api/products/", productcontroller.FindAllProductByBrandId)
	mux.HandleFunc("/api/purchase", purchasecontroller.Insert)
	mux.HandleFunc("/api/purchase/", purchasecontroller.FindById)

	return mux
}
