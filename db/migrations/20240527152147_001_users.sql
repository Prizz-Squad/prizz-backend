-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
                                     id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                     username VARCHAR UNIQUE,
                                     password VARCHAR,
                                     role INT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
