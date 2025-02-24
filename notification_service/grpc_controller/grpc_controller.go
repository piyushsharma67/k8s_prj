package grpc_controller

import (
	"notification_service/service"
)

type GrpcControllerStruct struct {
	service service.ServiceStruct
}

func InitialiseGrpcController(repository service.ServiceStruct) *GrpcControllerStruct {
	return &GrpcControllerStruct{
		service: repository,
	}
}
