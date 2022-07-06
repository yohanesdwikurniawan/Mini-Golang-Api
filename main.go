package main

import (
	"Catalys-Tech-Backend-Test/app"
	"Catalys-Tech-Backend-Test/controller/brandcontroller"
	"Catalys-Tech-Backend-Test/controller/ordercontroller"
	"Catalys-Tech-Backend-Test/controller/productcontroller"
	"Catalys-Tech-Backend-Test/helper"
	"Catalys-Tech-Backend-Test/repository/brandrepository"
	"Catalys-Tech-Backend-Test/repository/orderrepository"
	"Catalys-Tech-Backend-Test/repository/productrepository"
	"Catalys-Tech-Backend-Test/service/brandservice"
	"Catalys-Tech-Backend-Test/service/orderservice"
	"Catalys-Tech-Backend-Test/service/productservice"
	"embed"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	app.InitDatabase("root", "root", "localhost", "3306", "initDb")
	dbConnection := app.GetConnection("root", "root", "localhost", "3306", "initDb")
	app.RunMigrations(embedMigrations, dbConnection)

	validate := validator.New()

	brandrepository := brandrepository.NewBrandRepository()
	brandservice := brandservice.NewBrandService(brandrepository, dbConnection, validate)
	brandcontroller := brandcontroller.NewBrandController(brandservice)

	productrepository := productrepository.NewProductRepository()
	productservice := productservice.NewProductService(productrepository, dbConnection, validate)
	productcontroller := productcontroller.NewProductController(productservice)

	orderrepository := orderrepository.NewOrderRepository()
	orderservice := orderservice.NewOrderService(orderrepository, dbConnection, validate)
	ordercontroller := ordercontroller.NewOrderController(orderservice)

	router := app.AllRouter(brandcontroller, productcontroller, ordercontroller)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
