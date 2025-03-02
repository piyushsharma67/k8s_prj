package grpc_controller

import (
	"auth_service/controllers/common"
	"auth_service/services"
)

type GRPCController struct {
    *common.ControllerStruct
}

func (c *GRPCController)NewGRPCController(s *services.ServiceStruct) *GRPCController {
    c.ControllerStruct=&common.ControllerStruct{}
    c.ControllerStruct.InitialiseController(s)
    return c
}