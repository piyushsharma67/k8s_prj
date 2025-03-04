package services

import "notification_service/repository"

type ServiceStruct struct{
	repository *repository.Repository
}

func InitialiseService(repository *repository.Repository)*ServiceStruct{
	return &ServiceStruct{
		repository: repository,
	}
}