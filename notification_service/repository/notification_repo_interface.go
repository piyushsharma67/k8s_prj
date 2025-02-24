package repository

import "context"

type NotificationRepository interface{
	GetUserFcm(ctx context.Context,userId int32)
}