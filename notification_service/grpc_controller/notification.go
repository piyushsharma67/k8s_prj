package grpc_controller

import (
	"context"
	"fmt"
	"notification_service/proto"
)

func (c *GrpcControllerStruct)SendMobileNotification(ctx context.Context,details *proto.SendNotificationRequest)(*proto.SendNotificationResponse,error){
	fmt.Println("called")
	return nil,nil
}