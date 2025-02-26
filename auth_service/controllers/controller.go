package controllers

import (
	"auth_service/proto"
	"auth_service/services"
)


type ControllerStruct struct{
	proto.UnimplementedAuthServiceServer
	Service *services.ServiceStruct
}

func (c *ControllerStruct)InitialiseController(s *services.ServiceStruct)*ControllerStruct{
	c.Service=s

	return c
}