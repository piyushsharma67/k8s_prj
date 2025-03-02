package controllers

import (
	"auth_service/models"
	"auth_service/proto"
	"auth_service/utils"
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
)

func (c *ControllerStruct) SaveFcmInDb(ctx context.Context, details *proto.SaveUserFcmRequest) (*proto.SaveUserFcmResponse, error) {
	validate := validator.New()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	if err := validate.Struct(details); err != nil {
		return nil, utils.GRPCErrorResponse(codes.InvalidArgument, err.Error())
	}

	ctx, cancel := context.WithTimeout(ctx, utils.ApiTimeoutTime)

	defer cancel()

	_, err := c.Service.SaveUserFcm(ctx, &models.CreateUserFcm{
		UserId:   details.UserId,
		FcmToken: details.FcmToken,
	})

	if err != nil {
		return nil, err
	}

	return nil, nil

}

func (c *ControllerStruct) GetUserFcm(ctx context.Context, details *proto.GetUserFcmRequest) (*proto.GetUserFcmResponse, error) {
	validate := validator.New()

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	if err := validate.Struct(details); err != nil {
		return nil, utils.GRPCErrorResponse(codes.InvalidArgument, err.Error())
	}

	ctx, cancel := context.WithTimeout(ctx, utils.ApiTimeoutTime)

	defer cancel()

	userFcm, err := c.Service.GetUserFcm(ctx, details.UserId)

	if err != nil {
		return nil, err
	}

	return &proto.GetUserFcmResponse{
		FcmToken: userFcm.FcmToken,
		UserId: userFcm.UserID,
		Error: nil,
	}, nil

}
