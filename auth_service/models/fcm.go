package models

import "github.com/jackc/pgx/v5/pgtype"

type CreateUserFcm struct{
	UserId int32
	FcmToken string
}

type UserFcm struct {
	ID        int32            `json:"id" bson:"_id,omitempty"`
	UserID    int32            `json:"user_id" bson:"user_id"`
	FcmToken  string           `json:"fcm_token" bson:"fcm_token"`
	CreatedAt pgtype.Timestamp `json:"created_at" bson:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at" bson:"updated_at"`
}