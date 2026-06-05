package grpc_handler

import (
	"context"

	carv1 "github.com/course/tasks/api/car/v1"
)

type DashboardHandler struct {
	carv1.UnimplementedDashboardServiceServer
	// repo *db.Queries
}

// Task 6: CQRS (Read model with joins)
func (h *DashboardHandler) GetOwnerDashboard(ctx context.Context, req *carv1.GetOwnerDashboardRequest) (*carv1.GetOwnerDashboardResponse, error) {
	// data, err := h.repo.GetOwnerDashboard(ctx, req.OwnerId)
	
	// Example mapping from rich DB response
	return &carv1.GetOwnerDashboardResponse{
		OwnerId:   req.OwnerId,
		OwnerName: "John Doe",
		Cars: []*carv1.DashboardCar{
			{
				Vin:               "VIN123",
				Brand:             "Toyota",
				LastServiceDate:   "2023-10-01",
				ServiceDescription: "Oil change",
			},
		},
	}, nil
}
