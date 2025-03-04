package common

import "notification_service/services"


type ControllerStruct struct{
	Service *services.ServiceStruct
}

func (c *ControllerStruct)InitialiseController(s *services.ServiceStruct)*ControllerStruct{
	c.Service=s
	return c
}