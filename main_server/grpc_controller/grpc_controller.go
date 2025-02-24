package grpc_controller

import (
	"main_server/proto"
	"main_server/services"
)

type GrpcControllerStruct struct{
	proto.UnimplementedAuthServiceServer
	service *services.ServiceStruct
}

func (c *GrpcControllerStruct)InitialiseController(s *services.ServiceStruct)*GrpcControllerStruct{
	c.service=s
	return c
}