-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Purchases(
    Id int primary key auto_increment,
	ProductId          int NOT NULL,
	BrandId            int NOT NULL,
	ProductQty         int,
	UnitPrice          int,
	TotalPrice         int,
	CreatedAt          datetime default CURRENT_TIMESTAMP,
    FOREIGN KEY (ProductId) REFERENCES Products(Id),
    FOREIGN KEY (BrandId) REFERENCES Brands(Id)
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Purchases;
-- +goose StatementEnd
