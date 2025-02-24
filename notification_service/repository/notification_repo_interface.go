package repository

import "context"

type NotificationRepository interface{
	CreateUserFcmToken(ctx context.Context,userId int32,fcmToken string)
	GetUserFcm(ctx context.Context,userId int32)
}