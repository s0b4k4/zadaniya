package grpc_handler

import (
	"context"

	carv1 "github.com/course/tasks/api/car/v1"
)

type BookingHandler struct {
	carv1.UnimplementedBookingServiceServer
	// repo *db.Queries // with pgx transaction support
}

// Task 8: Locking
func (h *BookingHandler) CreateBooking(ctx context.Context, req *carv1.CreateBookingRequest) (*carv1.CreateBookingResponse, error) {
	// Example flow for pessimistic locking:
	// 1. tx, err := db.Begin()
	// 2. count, err := tx.CheckCarAvailability(ctx, CheckCarAvailabilityParams{...}) 
	//    This uses `FOR UPDATE` so concurrent requests will block until tx completes.
	// 3. if count > 0 { return Error(Not available) }
	// 4. tx.CreateBooking(ctx, ...)
	// 5. tx.Commit()
	
	return &carv1.CreateBookingResponse{
		BookingId: 1,
		Success:   true,
	}, nil
}

func (h *BookingHandler) CheckAvailability(ctx context.Context, req *carv1.CheckAvailabilityRequest) (*carv1.CheckAvailabilityResponse, error) {
	// Quick check without full lock
	return &carv1.CheckAvailabilityResponse{Available: true}, nil
}
