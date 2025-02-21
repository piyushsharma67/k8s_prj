-- database/queries/user.sql

-- Create user 
-- name: CreateUser :exec
-- params: CreateUserParams
-- returns: User
INSERT INTO users (name, email, password, created_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP)
RETURNING id, name, email, password, created_at;

-- Get user by email
-- name: GetUserByEmail :one
-- params: GetUserByEmailParams
-- returns: User
SELECT id, name, email, password, created_at
FROM users
WHERE email = $1;

-- name: GetUserAllPhotos :many
-- params: UserID
-- returns: UploadedFile
SELECT id, user_id, file_url, created_at
FROM uploaded_files
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;
