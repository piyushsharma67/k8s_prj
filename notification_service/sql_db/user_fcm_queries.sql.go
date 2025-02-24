// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: user_fcm_queries.sql

package sql_db

import (
	"context"
)

const createUserFcmToken = `-- name: CreateUserFcmToken :one

INSERT INTO user_fcm_tokens (user_id, fcm_token, created_at, updated_at)
VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, user_id, fcm_token, created_at, updated_at
`

type CreateUserFcmTokenParams struct {
	UserID   int32
	FcmToken string
}

// database/queries/user_fcm_token.sql
// params: CreateUserFcmTokenParams
// returns: UserFcmToken
func (q *Queries) CreateUserFcmToken(ctx context.Context, arg CreateUserFcmTokenParams) (UserFcmToken, error) {
	row := q.db.QueryRow(ctx, createUserFcmToken, arg.UserID, arg.FcmToken)
	var i UserFcmToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FcmToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserFcmTokenByUserID = `-- name: GetUserFcmTokenByUserID :one
SELECT id, user_id, fcm_token, created_at, updated_at
FROM user_fcm_tokens
WHERE user_id = $1
`

// params: GetUserFcmTokenByUserIDParams
// returns: UserFcmToken
func (q *Queries) GetUserFcmTokenByUserID(ctx context.Context, userID int32) (UserFcmToken, error) {
	row := q.db.QueryRow(ctx, getUserFcmTokenByUserID, userID)
	var i UserFcmToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FcmToken,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserFcmToken = `-- name: UpdateUserFcmToken :exec
UPDATE user_fcm_tokens
SET fcm_token = $2, updated_at = CURRENT_TIMESTAMP
WHERE user_id = $1
`

type UpdateUserFcmTokenParams struct {
	UserID   int32
	FcmToken string
}

// params: UpdateUserFcmTokenParams
func (q *Queries) UpdateUserFcmToken(ctx context.Context, arg UpdateUserFcmTokenParams) error {
	_, err := q.db.Exec(ctx, updateUserFcmToken, arg.UserID, arg.FcmToken)
	return err
}
