package cmd

import (
	"auth_service/enums"
	"auth_service/proto"
	"auth_service/repository"
	"auth_service/sql_db"
	"context"
	"fmt"
	configPkg "auth_service/config"

	"log"
	"net"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var port string
var dbType string
var connType string

func runGrpcServer(){
	addr := fmt.Sprintf(":%s", port)
	listener,err:=net.Listen("tcp",addr)
	if err!=nil{
		log.Fatal(err)
	}

	server:=grpc.NewServer()
	proto.RegisterAuthServiceServer(server,&proto.UnimplementedAuthServiceServer{})
	log.Println("Starting gRPC server on port", port)

	log.Printf("Starting gRPC server on port %s", port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}

func runHttpsServer(){
	log.Println("Starting https server on port", port)
}

var startServer = &cobra.Command{
	Use:  "start",
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

		if connType == string(enums.Grpc){
			runGrpcServer()
		}else{
			runHttpsServer()
		}

	},
}

func init() {
	startServer.Flags().StringVarP(&port, "port", "p", "8080", "Port to run the server on")
	startServer.Flags().StringVar(&dbType, "db", "postgres", "Database type (postgres or mongo)")
	startServer.Flags().StringVar(&connType, "conn", "grpc", "Connection type (grpc or https)")
	RootCommand.AddCommand(startServer)
}
