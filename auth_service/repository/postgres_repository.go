package repository

import "k8s_project/auth_service/sql_db"

type PostgresRepository struct {
	db *sql_db.Queries
}
func NewPostgresRepository(db *sql_db.Queries)*PostgresRepository{
	return &PostgresRepository{db:db}
}