package repository

import (
	"context"
	"errors"
	"k8s_project/auth_service/models"
)

// import (
// 	"context"
// 	"errors"
// 	"k8s_project/auth_service/models"
// )

func (r *PostgresRepository)CreateUser(ctx context.Context,userDetails models.CreateUserRequest)error{
	return errors.New("Hey i am called")
}

func (r *PostgresRepository)GetUserByuserId(ctx context.Context,userDetails models.GetUserByUserId)error{
	return errors.New("Hey i am called")
}
