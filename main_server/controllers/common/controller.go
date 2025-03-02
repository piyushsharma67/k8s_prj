package common

import (
	"main_server/proto"
	"main_server/services"
)

type ControllerStruct struct{
	Service *services.ServiceStruct
	AuthService proto.AuthServiceClient
}

func (c *ControllerStruct)InitialiseController(s *services.ServiceStruct,p proto.AuthServiceClient)*ControllerStruct{
	c.Service=s
	c.AuthService=p
	return c
}