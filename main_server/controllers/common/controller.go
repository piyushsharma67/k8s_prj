package common

import (
	"main_server/proto/auth"
	"main_server/proto/notification"
	"main_server/services"
)

type ControllerStruct struct {
	Service     *services.ServiceStruct
	AuthService auth.AuthServiceClient
	NotificationServcie notification.NotificationServiceClient
}

func (c *ControllerStruct) InitialiseController(s *services.ServiceStruct, p auth.AuthServiceClient,n  notification.NotificationServiceClient) *ControllerStruct {
	c.Service = s
	c.AuthService = p
	c.NotificationServcie=n
	return c
}
