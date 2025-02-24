package cmd

import (
	configPkg "auth_service/config"
	"auth_service/enums"
	"auth_service/grpc_controller"
	"auth_service/proto"
	"sync"

	"auth_service/repository"
	"auth_service/routes"
	"auth_service/services"
	"auth_service/sql_db"
	"context"
	"fmt"
	"net/http"

	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var port string
var dbType string
var connType string

func runGrpcServer(ctx context.Context, wg *sync.WaitGroup) {
	addr := fmt.Sprintf(":%s", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, &grpc_controller.GrpcControllerStruct{})
	log.Println("Starting gRPC server on port", port)

	go func() {
		log.Println("Starting gRPC server on port 50051...")
		if err := server.Serve(listener); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	<-ctx.Done()
	defer wg.Done()
}

func runHttpsServer(repo *repository.Repositories, ctx context.Context, wg *sync.WaitGroup) {
	service := services.ServiceStruct{}
	instance := service.InitialiseService(repo)
	r := routes.InitRoutes(instance)

	go func() {
		fmt.Println("Running http server on port", port)
		addr := fmt.Sprintf(":%s", port)
		if err := http.ListenAndServe(addr, r); err != nil {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	defer wg.Done()
}

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

		ctx, cancel := context.WithCancel(context.Background())
		var wg sync.WaitGroup
		wg.Add(2)

		go runGrpcServer(ctx, &wg)
		go runHttpsServer(repo,ctx,&wg)

		defer cancel()

	},
}

func init() {
	startServer.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
	startServer.Flags().StringVar(&dbType, "db", "postgres", "Database type (postgres or mongo)")
	startServer.Flags().StringVar(&connType, "conn", "grpc", "Connection type (grpc or https)")
	RootCommand.AddCommand(startServer)
}
