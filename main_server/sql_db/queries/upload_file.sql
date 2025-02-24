-- database/queries/upload_file.sql

-- name: CreateUploadedFile :exec
-- params: CreateUploadedFileParams
INSERT INTO uploaded_files (user_id, file_url, created_at, updated_at)
VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

