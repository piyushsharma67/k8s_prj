package grpc_controller

import (
	"notification_service/proto"
	"notification_service/service"
)

type GrpcControllerStruct struct {
	*proto.UnimplementedNotificationServiceServer
	Service *service.ServiceStruct
}

func InitialiseGrpcController(service *service.ServiceStruct) *GrpcControllerStruct {
	return &GrpcControllerStruct{
		Service: service,
	}
}


