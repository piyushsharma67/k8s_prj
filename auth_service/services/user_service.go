package services

import (
	"auth_service/models"
	"auth_service/sql_db"
	"auth_service/utils"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *ServiceStruct) InsertUserInDB(ctx context.Context, user *models.User) (*models.User, error) {
	
	_, err := r.Repository.AuthRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		// Check if error is "record not found"
		if errors.Is(err, pgx.ErrNoRows) {
			// User does not exist, so insert the user
			hashedPass, err := utils.HashPassword(user.Password)
			if err != nil {
				return nil, err
			}
			createUserParams := &sql_db.CreateUserParams{
				Name:     user.Name,
				Email:    user.Email,
				Password: user.Password,
			}
			createUserParams.Password = hashedPass
			err = r.Repository.AuthRepo.InsertUserInDB(ctx, *createUserParams)
			if err != nil {
				return nil, err
			}
			// Fetch newly created user
			newUser, err := r.Repository.AuthRepo.GetUserByEmail(ctx, user.Email)
			if err != nil {
				return nil, err
			}
			user.ID = newUser.ID
			token, err := utils.EncodeToken(user.ID)
			if err != nil {
				return nil, utils.INTERNAL_SERVER_ERROR
			}
			user.Token = token
			return user, nil
		}
		// Return error if it's not a "not found" error
		return nil, err
	}
	return nil, utils.USER_ALREADY_EXISTS
}

func (r *ServiceStruct) GetUserByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	user, err := r.Repository.AuthRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	return user,nil
}
