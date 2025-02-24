package grpc_controller

import (
	"notification_service/proto"
	"notification_service/service"
)

type GrpcControllerStruct struct {
	*proto.UnimplementedNotificationServiceServer
	service service.ServiceStruct
}

func InitialiseGrpcController(repository service.ServiceStruct) *GrpcControllerStruct {
	return &GrpcControllerStruct{
		service: repository,
	}
}


