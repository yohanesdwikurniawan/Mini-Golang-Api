-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS Brands(
    Id int primary key auto_increment, 
    Name text
); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE Brands;
-- +goose StatementEnd
