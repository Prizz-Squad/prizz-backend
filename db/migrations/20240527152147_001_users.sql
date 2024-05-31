-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT gen_random_uuid() PRIMARY key,
    Username VARCHAR(255) NOT NULL Unique,
    Password VARCHAR(255) NOT NULL,
    Department int not NULL,
    Role int not NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
