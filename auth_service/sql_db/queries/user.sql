-- database/queries/user.sql

-- Create user 
-- name: CreateUser :exec
-- params: CreateUserParams
-- returns: User
INSERT INTO users (name, email, password,is_active, created_at)
VALUES ($1, $2, $3,$4, CURRENT_TIMESTAMP)
RETURNING id, name, email, password,is_active, created_at;

-- Get user by email
-- name: GetUserByEmail :one
-- params: GetUserByEmailParams
-- returns: User
SELECT id, name, email, password,is_active, created_at
FROM users
WHERE email = $1;
