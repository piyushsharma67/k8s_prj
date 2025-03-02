package models

import (
	"main_server/proto"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserFcm struct {
	FcmToken string `json:"fcm_token"`
}

type UserFcmToken struct {
	ID        int32            `json:"id" bson:"_id,omitempty"`
	UserID    int32            `json:"user_id" bson:"user_id"`
	FcmToken  string           `json:"fcm_token" bson:"fcm_token"`
	CreatedAt pgtype.Timestamp `json:"created_at" bson:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at" bson:"updated_at"`
}

type SaveUserFcmRequestLowercase struct {
	FcmToken string `json:"fcmToken,omitempty"`
	UserId   int32  `json:"userId,omitempty"`
}

func ConvertToLowercaseRequest(req *proto.SaveUserFcmResponse) *SaveUserFcmRequestLowercase {
	return &SaveUserFcmRequestLowercase{
		FcmToken: req.GetFcmToken(),
		UserId:   req.GetUserId(),
	}
}