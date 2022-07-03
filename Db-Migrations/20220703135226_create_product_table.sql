-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Products(
    ProductId int primary key auto_increment, 
	ProductName  text,
	ProductPrice int,
    BrandId int,
	FOREIGN KEY (BrandId) REFERENCES Brands(BrandId)
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Products;
-- +goose StatementEnd
