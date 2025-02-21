package repository

import (
	"context"
	"k8s_project/auth_service/models"
)

type AuthRepository interface{
	CreateUser(context.Context,models.CreateUserRequest)error
	GetUserByuserId(context.Context,models.GetUserByUserId)error
} 