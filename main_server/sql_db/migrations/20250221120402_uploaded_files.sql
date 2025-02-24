-- +goose Up
-- +goose StatementBegin
CREATE TABLE uploaded_files(
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    file_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS uploaded_files;
-- +goose StatementEnd
