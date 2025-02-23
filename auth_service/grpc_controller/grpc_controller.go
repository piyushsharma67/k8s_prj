package grpc_controller

import (
	"auth_service/proto"
	"auth_service/services"
)

type GrpcControllerStruct struct{
	proto.UnimplementedAuthServiceServer
	service *services.ServiceStruct
}

func (c *GrpcControllerStruct)InitialiseController(s *services.ServiceStruct)*GrpcControllerStruct{
	c.service=s
	return c
}