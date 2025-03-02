package models

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	ID        int32     `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name" validate:"required,min=2"`
	Email     string    `json:"email" bson:"email" validate:"required,email"`
	Password  string    `json:"password,omitempty" bson:"password,omitempty" validate:"required,min=6"`
	Token     string    `json:"token,omitempty" bson:"token,omitempty"`
	CreatedAt pgtype.Timestamp `json:"created_at" bson:"created_at"`
}

type CreateUserRequest struct{
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type GetUserByUserId struct{
	Id int32 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

