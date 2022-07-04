-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Products(
    Id int primary key auto_increment, 
	Name  text,
	Price int,
    BrandId int,
	FOREIGN KEY (BrandId) REFERENCES Brands(Id)
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Products;
-- +goose StatementEnd
