package service

import "context"

func (s *ServiceStruct)GetFcmByUserid(ctx context.Context,userId int32){
	s.repository.Notification.GetUserFcm(ctx,userId)
}