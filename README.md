# Mini-Golang-Api

assumption
1. brandId and productId always focus on existing brands and products
2. price, unit price, and total price is always an integer
3. the unit price is always the same as the price data in the product table
4. total price is the sum of the quantity times the unit price

how to
1. clone or extract data to local device
2. get mysql driver, goose, and validator
    MySQL driver :  https://github.com/go-sql-driver/mysql/
    goose : https://github.com/pressly/goose
    validator : https://github.com/go-playground/validator
3. go to main.go file and set your own database user, password, hostname, port, and database name
4. run the main.go file and the project will automatically create new database if the database is not exist and run migration using goose, so you have to make sure that the mysql driver and goose are installed
5. Access endpoints using postman
