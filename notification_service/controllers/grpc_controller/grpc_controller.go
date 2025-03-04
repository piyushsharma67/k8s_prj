package grpc_controller

import (
	"notification_service/controllers/common"
	"notification_service/proto/auth"
	"notification_service/services"
)

type GRPCController struct {
    *common.ControllerStruct
    AuthClient auth.AuthServiceClient
}

func (c *GRPCController)NewGRPCController(s *services.ServiceStruct) *GRPCController {
    c.ControllerStruct=&common.ControllerStruct{}
    c.ControllerStruct.InitialiseController(s)
    return c
}