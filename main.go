package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// Generated
	// carv1 "github.com/course/tasks/api/car/v1"
	// documentv1 "github.com/course/tasks/api/document/v1"
	// warehousev1 "github.com/course/tasks/api/warehouse/v1"
	// smarthomev1 "github.com/course/tasks/api/smarthome/v1"
	// articlev1 "github.com/course/tasks/api/article/v1"
	// telemetryv1 "github.com/course/tasks/api/telemetry/v1"

	// Handlers
	// grpc_handler "github.com/course/tasks/internal/handler/grpc"
	// "github.com/course/tasks/internal/repository"
	// "github.com/course/tasks/internal/service"
)

func main() {
	log.Println("Starting Course Tasks Application...")

	// Dependency Injection
	// worker := service.NewWarehouseWorker()
	// go worker.Start(context.Background())
	// searchClient := service.NewSearchClient()
	// garageRepo := repository.NewGarageMemoryRepo()

	// Interceptors (Task 5)
	// interceptorOpt := grpc.UnaryInterceptor(grpc_handler.AuthInterceptor)
	// streamInterceptorOpt := grpc.StreamInterceptor(grpc_handler.AuthStreamInterceptor)

	grpcServer := grpc.NewServer(
		// interceptorOpt,
		// streamInterceptorOpt,
	)

	// Register Services
	// carv1.RegisterGarageServiceServer(grpcServer, grpc_handler.NewGarageHandler(garageRepo))
	// documentv1.RegisterDocumentServiceServer(grpcServer, &grpc_handler.DocumentHandler{})
	// carv1.RegisterConveyorServiceServer(grpcServer, &grpc_handler.ConveyorHandler{})
	// warehousev1.RegisterWarehouseServiceServer(grpcServer, grpc_handler.NewWarehouseHandler(worker))
	// smarthomev1.RegisterSmartHomeServiceServer(grpcServer, &grpc_handler.SmartHomeHandler{})
	// carv1.RegisterDashboardServiceServer(grpcServer, &grpc_handler.DashboardHandler{})
	// articlev1.RegisterContentServiceServer(grpcServer, grpc_handler.NewContentHandler(searchClient))
	// carv1.RegisterBookingServiceServer(grpcServer, &grpc_handler.BookingHandler{})
	// telemetryv1.RegisterTelemetryServiceServer(grpcServer, &grpc_handler.TelemetryHandler{})
	// carv1.RegisterMaintenanceServiceServer(grpcServer, &grpc_handler.MaintenanceHandler{})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Println("gRPC server listening on :50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// REST Gateway
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	// opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	
	// Register gateway handlers
	// carv1.RegisterGarageServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	// documentv1.RegisterDocumentServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	// carv1.RegisterConveyorServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	// warehousev1.RegisterWarehouseServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	// smarthomev1.RegisterSmartHomeServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	// carv1.RegisterDashboardServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	// articlev1.RegisterContentServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	// carv1.RegisterBookingServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	// telemetryv1.RegisterTelemetryServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)
	// carv1.RegisterMaintenanceServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts)

	log.Println("REST Gateway listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to serve REST gateway: %v", err)
	}
}
