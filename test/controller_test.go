package test

import (
	"Catalyst-Tech-Backend-Test/app"
	"Catalyst-Tech-Backend-Test/controller/brandcontroller"
	"Catalyst-Tech-Backend-Test/controller/ordercontroller"
	"Catalyst-Tech-Backend-Test/controller/productcontroller"
	"Catalyst-Tech-Backend-Test/repository/brandrepository"
	"Catalyst-Tech-Backend-Test/repository/orderrepository"
	"Catalyst-Tech-Backend-Test/repository/productrepository"
	"Catalyst-Tech-Backend-Test/service/brandservice"
	"Catalyst-Tech-Backend-Test/service/orderservice"
	"Catalyst-Tech-Backend-Test/service/productservice"
	"embed"
	"net/http"

	"github.com/go-playground/validator/v10"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func initSetupTest() *http.ServeMux {
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

	return app.AllRouter(brandcontroller, productcontroller, ordercontroller)
}