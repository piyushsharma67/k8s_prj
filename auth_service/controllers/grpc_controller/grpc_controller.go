package grpc_controller

import (
	"auth_service/controllers/common"
	"auth_service/services"
)

type GRPCController struct {
    *common.ControllerStruct
}

func NewGRPCController(s *services.ServiceStruct) *GRPCController {
    ctrl := &GRPCController{}
    ctrl.InitialiseController(s)
    return ctrl
}