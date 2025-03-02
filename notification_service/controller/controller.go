package controller

import (
	"notification_service/proto"
	"notification_service/service"
)


type ControllerStruct struct{
	proto.UnimplementedNotificationServiceServer
	Service *service.ServiceStruct
}

func (c *ControllerStruct)InitialiseController(service *service.ServiceStruct)*ControllerStruct{
	c.Service=service

	return c
}