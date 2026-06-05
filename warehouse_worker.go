package service

import (
	"context"
	"log"
	"time"

	"github.com/course/tasks/internal/domain"
	// "github.com/course/tasks/internal/repository/db"
)

type WarehouseWorker struct {
	orderChan chan domain.Order
}

func NewWarehouseWorker() *WarehouseWorker {
	return &WarehouseWorker{
		orderChan: make(chan domain.Order, 100),
	}
}

// Emulate putting to RabbitMQ
func (w *WarehouseWorker) PublishOrder(order domain.Order) {
	w.orderChan <- order
}

// Background goroutine reading the queue
func (w *WarehouseWorker) Start(ctx context.Context) {
	log.Println("Starting warehouse background worker...")
	for {
		select {
		case <-ctx.Done():
			log.Println("Stopping warehouse worker")
			return
		case order := <-w.orderChan:
			// Process order (e.g., reserve items)
			log.Printf("Processing order %d...\n", order.ID)
			time.Sleep(2 * time.Second) // Simulate work
			log.Printf("Order %d processed and status changed to COMPLETED\n", order.ID)
			
			// Here you would update DB using db.UpdateOrderStatus
		}
	}
}
