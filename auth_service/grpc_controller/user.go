package grpc_controller

import (
	"auth_service/proto"
	"context"
)

func (s *GrpcControllerStruct)Signup(ctx context.Context,details *proto.SignupRequest)(*proto.SignupResponse,error){
	return nil,nil
}

func (s *GrpcControllerStruct)Login(ctx context.Context,login *proto.LoginRequest)(*proto.LoginResponse,error){
	return nil,nil
}

func (s *GrpcControllerStruct)ValidateToken(context.Context, *proto.ValidateTokenRequest) (*proto.ValidateTokenResponse, error){
	return nil,nil
}