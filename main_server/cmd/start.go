package cmd

import (
	"context"
	"fmt"
	"log"
	configPkg "main_server/config"
	"main_server/enums"
	"main_server/repository"
	"main_server/routes"
	"main_server/services"
	"main_server/sql_db"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
)

var port string
var dbType string
var connType string

var startServer = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		if env == "" {
			log.Fatal("Env not provided")
		}
		configPrj, err := configPkg.LoadConfig(env)
		if err != nil {
			log.Fatal(err)
		}
		var repo *repository.Repositories
		if dbType == string(enums.Postgres) {
			pgxpool, err := pgxpool.New(context.Background(), configPrj.GetDSN())

			if err != nil {
				log.Fatal(err)
			}
			queries := sql_db.New(pgxpool)
			repo, err = repository.InitialiseRepositories(enums.DBType(dbType), queries, nil)

			if err != nil {
				log.Fatal(err)
			}
		}

		service := services.ServiceStruct{}
		instance := service.InitialiseService(repo)
		r := routes.InitRoutes(instance)
		fmt.Println("env is",env)
		if env!="local"{
			port=os.Getenv("HTTP_PORT")
		}

		fmt.Println("Running http server on port", port)
		addr := fmt.Sprintf(":%s", port)
		if err := http.ListenAndServe(addr, r); err != nil {
			log.Fatal(err)
		}

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		<-stop // Wait for termination signal
		log.Println("Shutting down services...")

	},
}

func init() {
	startServer.Flags().StringVarP(&port, "port", "p", "3000", "Port to run the server on")
	startServer.Flags().StringVar(&dbType, "db", "postgres", "Database type (postgres or mongo)")
	startServer.Flags().StringVar(&connType, "conn", "grpc", "Connection type (grpc or https)")
	RootCommand.AddCommand(startServer)
}
