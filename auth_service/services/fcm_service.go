package services

import (
	"auth_service/models"
	"context"
)

func (s *ServiceStruct) SaveUserFcm(ctx context.Context, details *models.CreateUserFcm) (*models.UserFcm, error) {
	userDetails, err := s.Repository.AuthRepo.InsertUserFcmById(ctx, &models.CreateUserFcm{
		UserId:   details.UserId,
		FcmToken: details.FcmToken,
	})

	if err != nil {
		return nil, err
	}

	return userDetails, err
}

func (s *ServiceStruct) GetUserFcm(ctx context.Context, details *models.CreateUserFcm) (*models.UserFcm, error) {
	userDetails, err := s.Repository.AuthRepo.GetUserFcmById(ctx, details.UserId)

	if err != nil {
		return nil, err
	}

	return &models.UserFcm{
		FcmToken: userDetails.FcmToken,
		UserID:   userDetails.UserID,
	}, nil
}
