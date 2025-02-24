package grpc_controller

import (
	"context"
	"fmt"
	"notification_service/proto"
)

func (c *GrpcControllerStruct)SendMobileNotification(ctx context.Context,details *proto.SendNotificationRequest){
	fmt.Println("called")
}