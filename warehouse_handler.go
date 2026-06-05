package grpc_handler

import (
	"context"
	"math/rand"

	warehousev1 "github.com/course/tasks/api/warehouse/v1"
	"github.com/course/tasks/internal/domain"
	"github.com/course/tasks/internal/service"
)

type WarehouseHandler struct {
	warehousev1.UnimplementedWarehouseServiceServer
	worker *service.WarehouseWorker
}

func NewWarehouseHandler(worker *service.WarehouseWorker) *WarehouseHandler {
	return &WarehouseHandler{worker: worker}
}

func (h *WarehouseHandler) CreateOrder(ctx context.Context, req *warehousev1.CreateOrderRequest) (*warehousev1.CreateOrderResponse, error) {
	// 1. Create order in DB with status "PENDING"
	orderID := rand.Int31()

	// 2. Publish to message broker (RabbitMQ/Channel)
	h.worker.PublishOrder(domain.Order{
		ID:        int(orderID),
		ProductID: int(req.ProductId),
		Quantity:  int(req.Quantity),
		Status:    "PENDING",
	})

	// 3. Return immediately (Async processing)
	return &warehousev1.CreateOrderResponse{
		OrderId: orderID,
		Status:  "PENDING",
	}, nil
}

func (h *WarehouseHandler) GetOrderStatus(ctx context.Context, req *warehousev1.GetOrderStatusRequest) (*warehousev1.GetOrderStatusResponse, error) {
	// status := h.repo.GetOrderStatus(req.OrderId)
	return &warehousev1.GetOrderStatusResponse{
		OrderId: req.OrderId,
		Status:  "COMPLETED", // Example
	}, nil
}
