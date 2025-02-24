package repository

import (
	"notification_service/enums"
	"notification_service/sql_db"
	"notification_service/utils"

	"go.mongodb.org/mongo-driver/mongo"
)


type Repository struct{
	Notification NotificationRepository
}

func InitialiseRepositories(dbType enums.DBType, postgresDB *sql_db.Queries, mongoClient *mongo.Client) (*Repository, error) {

	if postgresDB == nil && mongoClient == nil {
		return nil, utils.DB_INSTANCE_REQUIRED
	}
	var notificationRepo NotificationRepository

	switch dbType {
	case enums.Postgres:
		notificationRepo = InitialiseSqlRepository(postgresDB)
	}

	return &Repository{
		Notification: notificationRepo,
	}, nil
}