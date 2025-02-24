package grpc_controller

import (
	"main_server/models"
	"main_server/proto"
	"main_server/utils"
	"context"

	"github.com/go-playground/validator"
	"google.golang.org/grpc/codes"
)

func (s *GrpcControllerStruct)Signup(ctx context.Context,details *proto.SignupRequest)(*proto.SignupResponse,error){
	validate := validator.New()
	if err := validate.Struct(details); err != nil {
		return nil,utils.GRPCErrorResponse(codes.InvalidArgument,err.Error())
	}

	ctx, cancel := context.WithTimeout(ctx, utils.ApiTimeoutTime)

	defer cancel()
	
	user,err:=s.service.InsertUserInDB(ctx,&models.User{
		Email: details.Email,
		Name: details.Name,
		Password: details.Password,
	})

	if err!=nil{
		return nil,utils.GRPCErrorResponse(codes.InvalidArgument,err.Error())
	}
	token,err:=utils.EncodeToken(user.ID)

	if err!=nil{
		return nil,utils.GRPCErrorResponse(codes.InvalidArgument,err.Error())
	}
	return &proto.SignupResponse{
		Name: user.Name,
		Email: user.Email,
		Id: user.ID,
		Token: token,
	},nil
}

func (s *GrpcControllerStruct)Login(ctx context.Context,login *proto.LoginRequest)(*proto.LoginResponse,error){
	validate := validator.New()
	if err := validate.Struct(login); err != nil {
		return nil,utils.GRPCErrorResponse(codes.InvalidArgument,err.Error())
	}

	ctx, cancel := context.WithTimeout(ctx, utils.ApiTimeoutTime)

	defer cancel()
	
	user,err:=s.service.GetUserByEmail(ctx,&models.User{
		Email: login.Email,
		Password: login.Password,
	})

	if err!=nil{
		return nil,utils.GRPCErrorResponse(codes.InvalidArgument,err.Error())
	}
	token,err:=utils.EncodeToken(user.ID)

	if err!=nil{
		return nil,utils.GRPCErrorResponse(codes.InvalidArgument,err.Error())
	}
	return &proto.LoginResponse{
		Name: user.Name,
		Email: user.Email,
		Id: user.ID,
		Token: token,
	},nil
}

func (s *GrpcControllerStruct)ValidateToken(context.Context, *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error){
	return nil,nil
}