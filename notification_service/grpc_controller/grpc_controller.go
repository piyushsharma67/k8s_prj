package grpc_controller

import "notification_service/repository"


type GrpcControllerStruct struct{
	repository repository.Repository
}

func InitialiseGrpcController(repository repository.Repository)*GrpcControllerStruct{
	return &GrpcControllerStruct{
		repository: repository,
	}
}