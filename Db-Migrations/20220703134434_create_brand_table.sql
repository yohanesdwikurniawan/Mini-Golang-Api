-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Brands(
    BrandId int primary key auto_increment, 
    BrandName text
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Brands;
-- +goose StatementEnd
