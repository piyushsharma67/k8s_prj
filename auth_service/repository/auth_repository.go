package repository

import (
	"context"
	"auth_service/models"
	"auth_service/sql_db"
)

func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := r.db.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Password:  user.Password,
		Token:     "",
		CreatedAt: user.CreatedAt,
	}, nil
}

func (r *PostgresRepository) InsertUserInDB(ctx context.Context, details sql_db.CreateUserParams) error {
	return r.db.CreateUser(ctx, details)
}

func (r *PostgresRepository) GetUserFcmById(ctx context.Context, userId int32) (*models.UserFcm, error) {
	userFcm, err := r.db.GetUserFcmTokenByUserID(ctx, userId)

	if err != nil {
		return nil, err
	}
	return &models.UserFcm{
		ID:        userFcm.ID,
		UserID:    userFcm.UserID,
		FcmToken:  userFcm.FcmToken,
		CreatedAt: userFcm.CreatedAt,
		UpdatedAt: userFcm.UpdatedAt,
	}, nil
}

func (r *PostgresRepository) InsertUserFcmById(ctx context.Context, details *models.CreateUserFcm) (*models.UserFcm, error) {
	userFcm, err := r.db.CreateUserFcmToken(ctx, sql_db.CreateUserFcmTokenParams{
		UserID: details.UserId,
		FcmToken: details.FcmToken,
	})

	if err != nil {
		return nil, err
	}
	return &models.UserFcm{
		ID:        userFcm.ID,
		UserID:    userFcm.UserID,
		FcmToken:  userFcm.FcmToken,
		CreatedAt: userFcm.CreatedAt,
		UpdatedAt: userFcm.UpdatedAt,
	}, nil
}

func (r *PostgresRepository) UpdateUserFcmById(ctx context.Context, details *models.CreateUserFcm) error {
	return r.db.UpdateUserFcmToken(ctx, sql_db.UpdateUserFcmTokenParams{
		UserID: details.UserId,
		FcmToken: details.FcmToken,
	})
}
