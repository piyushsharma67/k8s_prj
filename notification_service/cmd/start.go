package cmd

import (
	"context"
	"log"
	"net"
	"net/http"
	"notification_service/proto"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func startGrpcServer(ctx context.Context,wg *sync.WaitGroup) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Register your gRPC service here
	proto.RegisterNotificationServiceServer(grpcServer, &proto.UnimplementedNotificationServiceServer{})

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
func startHttpServer(ctx context.Context,wg *sync.WaitGroup) {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	server := &http.Server{Addr: ":8080", Handler: mux}

	go func() {
		log.Println("Starting HTTP server on port 8080...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	<-ctx.Done() // Wait for shutdown signal
	log.Println("Shutting down HTTP server...")

	// Graceful shutdown
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("HTTP server shutdown failed: %v", err)
	}

	defer wg.Done()
}

// Function to start SQS listener with graceful shutdown
func startSqsListener(ctx context.Context,wg *sync.WaitGroup) {
	go func() {
		log.Println("Starting SQS listener...")
		ticker := time.NewTicker(2 * time.Second) 

		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				log.Println("Shutting down SQS listener...")
				return
			case <-ticker.C: // This will run every time the ticker ticks
			log.Println("Listening for SQS messages...")
			}
		}
	}()

	defer wg.Done()
}
var startServer = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())

		var wg sync.WaitGroup

		wg.Add(3) // Number of services to wait for

		go startGrpcServer(ctx, &wg)
		go startHttpServer(ctx, &wg)
		go startSqsListener(ctx, &wg)

		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

		<-stop // Wait for termination signal
		log.Println("Shutting down services...")
		cancel() // Notify all services to stop

		wg.Wait() // Ensure all services shut down before exiting
		log.Println("All services stopped successfully")
	},
}

func init(){
	RootCommand.AddCommand(startServer)
}