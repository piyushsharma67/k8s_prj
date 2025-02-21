package cmd

import (
	"context"
	configPkg "k8s_project/auth_service/config"
	"k8s_project/auth_service/enums"
	"k8s_project/auth_service/repository"
	"k8s_project/auth_service/sql_db"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
)

var port string
var dbType string

var startServer = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		if env == "" {
			log.Fatal("Env not provided")
		}
		configPrj, err := configPkg.Loadconfig(env)
		if err != nil {
			log.Fatal(err)
		}

		if dbType == string(enums.Postgres) {
			pgxpool, err := pgxpool.New(context.Background(), configPrj.GetDSN())

			if err != nil {
				log.Fatal(err)
			}
			queries := sql_db.New(pgxpool)
			_, err = repository.InitialiseRepositories(enums.DBType(dbType), queries, nil)

			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	startServer.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
	startServer.Flags().StringVar(&dbType, "db-type", "postgres", "Database type (postgres or mongo)")
	RootCommand.AddCommand(startServer)
}
