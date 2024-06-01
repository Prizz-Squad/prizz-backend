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


-- triggering function that calculates the time
CREATE OR REPLACE FUNCTION calculate_hours_trigger()
RETURNS TRIGGER AS $$
BEGIN
    NEW.hours := EXTRACT(EPOCH FROM NEW.end_date - NEW.start_date) / 3600;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create the trigger
CREATE TRIGGER calculate_hours_trigger_report
    BEFORE INSERT ON report
    FOR EACH ROW
    EXECUTE FUNCTION calculate_hours_trigger();