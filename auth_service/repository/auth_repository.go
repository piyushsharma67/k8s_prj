package repository

import (
	"context"
	"k8s_project/auth_service/models"
	"k8s_project/auth_service/sql_db"
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

func (r *PostgresRepository) GetUserFcmById(ctx context.Context, userId int32) (*models.UserFcmToken, error) {
	userFcm, err := r.db.GetUserFcmTokenByUserID(ctx, userId)

	if err != nil {
		return nil, err
	}
	return &models.UserFcmToken{
		ID:        userFcm.ID,
		UserID:    userFcm.UserID,
		FcmToken:  userFcm.FcmToken,
		CreatedAt: userFcm.CreatedAt,
		UpdatedAt: userFcm.UpdatedAt,
	}, nil
}

func (r *PostgresRepository) InsertUserFcmById(ctx context.Context, details sql_db.CreateUserFcmTokenParams) (*models.UserFcmToken, error) {
	userFcm, err := r.db.CreateUserFcmToken(ctx, details)

	if err != nil {
		return nil, err
	}
	return &models.UserFcmToken{
		ID:        userFcm.ID,
		UserID:    userFcm.UserID,
		FcmToken:  userFcm.FcmToken,
		CreatedAt: userFcm.CreatedAt,
		UpdatedAt: userFcm.UpdatedAt,
	}, nil
}

func (r *PostgresRepository) UpdateUserFcmById(ctx context.Context, details sql_db.UpdateUserFcmTokenParams) error {
	return r.db.UpdateUserFcmToken(ctx, details)
}
