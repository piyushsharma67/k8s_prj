package controller

import (
	"context"
	"fmt"
	"notification_service/proto"
	"notification_service/utils"

	"github.com/go-playground/validator"
	"google.golang.org/grpc/codes"
)

func (c *ControllerStruct)SaveFcmGrpc(ctx context.Context,details *proto.SendNotificationRequest)(*proto.SendNotificationResponse,error){
    validate := validator.New()

	defer  func(){
		if r:=recover();r!=nil{
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	if err := validate.Struct(details); err != nil {
		return nil, utils.GRPCErrorResponse(codes.InvalidArgument, err.Error())
	}

	ctx, cancel := context.WithTimeout(ctx, utils.ApiTimeoutTime)

	defer cancel()


    return nil,nil
}