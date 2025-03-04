package http_controller

import (
	"notification_service/controllers/common"
	"notification_service/services"
)


type HTTPController struct {
    *common.ControllerStruct
    Service *services.ServiceStruct
}

func (c *HTTPController)NewHTTPController(s *services.ServiceStruct) *HTTPController {
    c.ControllerStruct = &common.ControllerStruct{}
    c.ControllerStruct.InitialiseController(s)

    return c
}