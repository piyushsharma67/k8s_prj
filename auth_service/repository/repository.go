package repository

import (
	"auth_service/enums"
	"auth_service/sql_db"
	"auth_service/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repositories struct {
	AuthRepo AuthRepository
}

func InitialiseRepositories(dbType enums.DBType, postgresDB *sql_db.Queries, mongoClient *mongo.Client) (*Repositories, error) {

	if postgresDB == nil && mongoClient == nil {
		return nil, utils.DB_INSTANCE_REQUIRED
	}
	var authRepo AuthRepository

	switch dbType {
	case enums.Postgres:
		authRepo = NewPostgresRepository(postgresDB)
	}

	return &Repositories{
		AuthRepo: authRepo,
	}, nil
}
