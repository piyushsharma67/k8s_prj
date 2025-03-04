package cmd

import (
	"context"
	"log"
	"net"
	configPkg "notification_service/config"
	"notification_service/controllers/grpc_controller"
	"notification_service/enums"
	"notification_service/proto/auth"
	"notification_service/proto/notification"
	"notification_service/repository"
	"notification_service/services"
	"notification_service/sql_db"

	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var port string
var dbType string
var connType string

func startGrpcServer(ctx context.Context, wg *sync.WaitGroup, repo *repository.Repository) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	service := services.InitialiseService(repo)
	grpcController := &grpc_controller.GRPCController{}

	controller := grpcController.NewGRPCController(service)

	// Register your gRPC service here
	notification.RegisterNotificationServiceServer(grpcServer, controller)
	auth.RegisterAuthServiceServer(grpcServer, auth.UnimplementedAuthServiceServer{})

	go func() {
		log.Println("Starting gRPC server on port 50051...")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("gRPC server error: %v", err)
		}
	}()

	<-ctx.Done() // Wait for shutdown signal
	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()

	defer wg.Done()
}

var startServer = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())

		if env == "" {
			log.Fatal("Env not provided")
		}
		configPrj, err := configPkg.LoadConfig(env)
		if err != nil {
			log.Fatal(err)
		}
		var repo *repository.Repository
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

		var wg sync.WaitGroup

		wg.Add(1) // Number of services to wait for

		go startGrpcServer(ctx, &wg, repo)
		// go startHttpServer(ctx, &wg)
		// go startSqsListener(ctx, &wg)

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		<-stop // Wait for termination signal
		log.Println("Shutting down services...")
		cancel() // Notify all services to stop

		wg.Wait() // Ensure all services shut down before exiting
		log.Println("All services stopped successfully")
	},
}

func init() {
	RootCommand.AddCommand(startServer)
}
