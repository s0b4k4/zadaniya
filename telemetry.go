package grpc_handler

import (
	"io"
	"log"

	// telemetryv1 "github.com/course/tasks/api/telemetry/v1"
)

type TelemetryHandler struct {
	// telemetryv1.UnimplementedTelemetryServiceServer
}

// Emulating the generated grpc code for bidirectional stream
type TelemetryStream interface {
	Recv() (interface{}, error) // Returns *telemetryv1.Signal
	Send(interface{}) error     // Accepts *telemetryv1.Alert
}

// Task 9: Bidirectional Stream
func (h *TelemetryHandler) TelemetryStream(stream TelemetryStream) error {
	for {
		// 1. Receive stream of signals
		signal, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("Received signal: %+v\n", signal)

		// Type assertion for demonstration
		// sig := signal.(*telemetryv1.Signal)

		// 2. High load processing logic...
		// if sig.Speed > 150 || sig.Rpm > 6000 {
			
		// 3. Send stream of alerts back
		err = stream.Send(nil /* &telemetryv1.Alert{...} */)
		if err != nil {
			return err
		}
		// }
	}
}
