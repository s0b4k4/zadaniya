package grpc_handler

import (
	"context"
	"time"

	smarthomev1 "github.com/course/tasks/api/smarthome/v1"
)

type SmartHomeHandler struct {
	smarthomev1.UnimplementedSmartHomeServiceServer
}

func (h *SmartHomeHandler) SubmitReading(ctx context.Context, req *smarthomev1.SubmitReadingRequest) (*smarthomev1.SubmitReadingResponse, error) {
	// Save to DB
	return &smarthomev1.SubmitReadingResponse{Success: true}, nil
}

// Task 5: Server Streaming
func (h *SmartHomeHandler) MonitorReadings(req *smarthomev1.MonitorReadingsRequest, stream smarthomev1.SmartHomeService_MonitorReadingsServer) error {
	// Simulate sending a stream of readings
	for i := 0; i < 10; i++ {
		err := stream.Send(&smarthomev1.Reading{
			DeviceId:  req.DeviceId,
			Value:     float32(20.0 + float32(i)),
			Timestamp: time.Now().Format(time.RFC3339),
		})
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second) // Simulate delay between readings
	}
	return nil
}
