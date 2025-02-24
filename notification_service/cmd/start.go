package cmd

import (
	"context"
	"log"
	"net"
	configPkg "notification_service/config"
	"notification_service/enums"
	"notification_service/grpc_controller"
	"notification_service/proto"
	"notification_service/repository"
	"notification_service/service"
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

func startGrpcServer(ctx context.Context, wg *sync.WaitGroup,repo *repository.Repository) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	service:=service.InitialiseService(repo)
	grpcController:=&grpc_controller.GrpcControllerStruct{
		Service:service,
	}

	// Register your gRPC service here
	proto.RegisterNotificationServiceServer(grpcServer, grpcController)

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

// Function to start HTTP server with graceful shutdown
// func startHttpServer(ctx context.Context,wg *sync.WaitGroup) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("OK"))
// 	})

// 	server := &http.Server{Addr: ":8080", Handler: mux}

// 	go func() {
// 		log.Println("Starting HTTP server on port 8080...")
// 		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			log.Fatalf("HTTP server error: %v", err)
// 		}
// 	}()

// 	<-ctx.Done() // Wait for shutdown signal
// 	log.Println("Shutting down HTTP server...")

// 	// Graceful shutdown
// 	if err := server.Shutdown(context.Background()); err != nil {
// 		log.Fatalf("HTTP server shutdown failed: %v", err)
// 	}

// 	defer wg.Done()
// }

// // Function to start SQS listener with graceful shutdown
// func startSqsListener(ctx context.Context,wg *sync.WaitGroup) {
// 	go func() {
// 		log.Println("Starting SQS listener...")
// 		ticker := time.NewTicker(2 * time.Second)

// 		defer ticker.Stop()
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				log.Println("Shutting down SQS listener...")
// 				return
// 			case <-ticker.C: // This will run every time the ticker ticks
// 			log.Println("Listening for SQS messages...")
// 			}
// 		}
// 	}()

//		defer wg.Done()
//	}
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

		go startGrpcServer(ctx, &wg,repo)
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
