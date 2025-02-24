-- database/queries/user_fcm_token.sql

-- name: CreateUserFcmToken :one
-- params: CreateUserFcmTokenParams
-- returns: UserFcmToken
INSERT INTO user_fcm_tokens (user_id, fcm_token, created_at, updated_at)
VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, user_id, fcm_token, created_at, updated_at;


-- name: GetUserFcmTokenByUserID :one
-- params: GetUserFcmTokenByUserIDParams
-- returns: UserFcmToken
SELECT id, user_id, fcm_token, created_at, updated_at
FROM user_fcm_tokens
WHERE user_id = $1;

-- name: UpdateUserFcmToken :exec
-- params: UpdateUserFcmTokenParams
UPDATE user_fcm_tokens
SET fcm_token = $2, updated_at = CURRENT_TIMESTAMP
WHERE user_id = $1;
