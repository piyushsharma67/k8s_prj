-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE users
ADD COLUMN is_active BOOLEAN DEFAULT true;  -- Adding the 'is_active' column

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE users
DROP COLUMN IF EXISTS is_active;  -- Removing the 'is_active' column

-- +goose StatementEnd
