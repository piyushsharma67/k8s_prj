package repository

import (
	"context"
	"notification_service/sql_db"
)


type SqlRepositoryStruct struct{
	db *sql_db.Queries
}

func InitialiseSqlRepository(db *sql_db.Queries)*SqlRepositoryStruct{
	return &SqlRepositoryStruct{
		db: db,
	}
}

func (r *SqlRepositoryStruct)CreateUserFcmToken(ctx context.Context,userId int32,fcmToken string){
	r.db.CreateUserFcmToken(ctx,sql_db.CreateUserFcmTokenParams{
		UserID: userId,
		FcmToken: fcmToken,
	})
}

func (r *SqlRepositoryStruct)GetUserFcm(ctx context.Context,userId int32){
	r.db.GetUserFcmTokenByUserID(ctx,userId)
}