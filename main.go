package main

import (
	"Catalys-Tech-Backend-Test/app"
	"Catalys-Tech-Backend-Test/controller/brandcontroller"
	"Catalys-Tech-Backend-Test/controller/productcontroller"
	"Catalys-Tech-Backend-Test/controller/purchasecontroller"
	"Catalys-Tech-Backend-Test/helper"
	"Catalys-Tech-Backend-Test/repository/brandrepository"
	"Catalys-Tech-Backend-Test/repository/productrepository"
	"Catalys-Tech-Backend-Test/repository/purchaserepository"
	"Catalys-Tech-Backend-Test/service/brandservice"
	"Catalys-Tech-Backend-Test/service/productservice"
	"Catalys-Tech-Backend-Test/service/purchaseservice"
	"embed"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	app.InitDatabase("initDb")
	dbConnection := app.GetConnection("initDb")
	app.RunMigrations(embedMigrations, dbConnection)

	validate := validator.New()

	brandrepository := brandrepository.NewBrandRepository()
	brandservice := brandservice.NewBrandService(brandrepository, dbConnection, validate)
	brandcontroller := brandcontroller.NewBrandController(brandservice)

	productrepository := productrepository.NewProductRepository()
	productservice := productservice.NewProductService(productrepository, dbConnection, validate)
	productcontroller := productcontroller.NewProductController(productservice)

	purchaserepository := purchaserepository.NewPurchaseRepository()
	purchaseservice := purchaseservice.NewPurchaseService(purchaserepository, dbConnection, validate)
	purchasecontroller := purchasecontroller.NewPurchaseController(purchaseservice)

	router := app.AllRouter(brandcontroller, productcontroller, purchasecontroller)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
