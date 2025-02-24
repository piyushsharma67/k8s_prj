package repository

import (
	"context"
	"main_server/models"
	"main_server/sql_db"
)

type AuthRepository interface{
	GetUserByEmail(ctx context.Context,email string)(*models.User,error)
	InsertUserInDB(ctx context.Context,details sql_db.CreateUserParams)error
	GetUserFcmById(ctx context.Context,userId int32)(*models.UserFcmToken,error)
	InsertUserFcmById(ctx context.Context,details sql_db.CreateUserFcmTokenParams)(*models.UserFcmToken,error)
	UpdateUserFcmById(ctx context.Context,details sql_db.UpdateUserFcmTokenParams)(error)
} 