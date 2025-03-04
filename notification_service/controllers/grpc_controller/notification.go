package grpc_controller

import (
	"context"
	"notification_service/proto/notification"
)

func (c *GRPCController)SendPushNotification(ctx context.Context,details *notification.SendPushNotificationRequest)(*notification.SendPushNotificationRequest,error){
	return nil,nil
}