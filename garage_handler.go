package grpc_handler

import (
	"context"

	carv1 "github.com/course/tasks/api/car/v1"
	"github.com/course/tasks/internal/domain"
	"github.com/course/tasks/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GarageHandler struct {
	carv1.UnimplementedGarageServiceServer
	repo *repository.GarageMemoryRepo
}

func NewGarageHandler(repo *repository.GarageMemoryRepo) *GarageHandler {
	return &GarageHandler{
		repo: repo,
	}
}

func (h *GarageHandler) CreateCar(ctx context.Context, req *carv1.CreateCarRequest) (*carv1.CarResponse, error) {
	if req.Vin == "" {
		return nil, status.Error(codes.InvalidArgument, "VIN is required")
	}

	car := domain.Car{
		VIN:   req.Vin,
		Brand: req.Brand,
		Year:  int(req.Year),
	}

	err := h.repo.CreateCar(ctx, car)
	if err != nil {
		if err.Error() == "car already exists" {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &carv1.CarResponse{
		Car: &carv1.Car{
			Vin:   car.VIN,
			Brand: car.Brand,
			Year:  int32(car.Year),
		},
	}, nil
}

func (h *GarageHandler) GetCar(ctx context.Context, req *carv1.GetCarRequest) (*carv1.CarResponse, error) {
	if req.Vin == "" {
		return nil, status.Error(codes.InvalidArgument, "VIN is required")
	}

	car, err := h.repo.GetCar(ctx, req.Vin)
	if err != nil {
		if err.Error() == "car not found" {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &carv1.CarResponse{
		Car: &carv1.Car{
			Vin:   car.VIN,
			Brand: car.Brand,
			Year:  int32(car.Year),
		},
	}, nil
}
