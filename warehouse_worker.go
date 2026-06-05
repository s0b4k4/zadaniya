package service

import (
    "fmt"
    "time"
)

// Task 4: Background Worker
type WarehouseWorker struct {
    // Usually holds a channel or connection to RabbitMQ
    msgQueue <-chan string
}

func (w *WarehouseWorker) Run() {
    for msg := range w.msgQueue {
        // Imitate background processing
        fmt.Printf("Processing order: %s\n", msg)
        time.Sleep(2 * time.Second)
        fmt.Printf("Order %s processed successfully\n", msg)
    }
}
