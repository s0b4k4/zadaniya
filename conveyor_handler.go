package grpc_handler

import (
	"context"

	carv1 "github.com/course/tasks/api/car/v1"
)

type ConveyorHandler struct {
	carv1.UnimplementedConveyorServiceServer
	// repo *db.Queries
}

func (h *ConveyorHandler) AssembleCar(ctx context.Context, req *carv1.AssembleCarRequest) (*carv1.AssembleCarResponse, error) {
	// err := h.repo.AssembleCar(ctx, db.AssembleCarParams{
	// 	EngineID:       req.EngineId,
	// 	TransmissionID: req.TransmissionId,
	// 	Vin:            req.Vin,
	// })
	return &carv1.AssembleCarResponse{Success: true}, nil
}

func (h *ConveyorHandler) GetCarSpec(ctx context.Context, req *carv1.GetCarSpecRequest) (*carv1.GetCarSpecResponse, error) {
	// spec, err := h.repo.GetCarSpec(ctx, req.Vin)
	return &carv1.GetCarSpecResponse{
		Car: &carv1.Car{Vin: req.Vin}, // Populate from spec
		Engine: &carv1.EngineSpec{
			Model:      "V8", // spec.EngineModel
			Horsepower: 500,  // spec.EngineHp
		},
		Transmission: &carv1.TransmissionSpec{
			Type: "Automatic", // spec.TransmissionType
		},
	}, nil
}
