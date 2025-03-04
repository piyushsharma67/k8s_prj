package common

import (
	"notification_service/proto/notification"
	"notification_service/services"
)


type ControllerStruct struct{
	notification.UnimplementedNotificationServiceServer
	Service *services.ServiceStruct
}

func (c *ControllerStruct)InitialiseController(s *services.ServiceStruct)*ControllerStruct{
	c.Service=s
	return c
}