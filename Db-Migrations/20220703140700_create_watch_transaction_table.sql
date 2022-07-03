-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS WatchTransactions(
    WatchTransactionId int primary key auto_increment,
	ProductId          int NOT NULL,
	BrandId            int NOT NULL,
	ProductQty         int,
	UnitPrice          int,
	TotalPrice         int,
	CreatedAt          datetime default CURRENT_TIMESTAMP,
    FOREIGN KEY (ProductId) REFERENCES Products(ProductId),
    FOREIGN KEY (BrandId) REFERENCES Brands(BrandId)
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE WatchTransactions;
-- +goose StatementEnd
