package grpc_handler

import (
    "io"
    // "github.com/course/tasks/pkg/telemetry_v1"
)

// Task 9: Bidirectional Streaming
// Imagine pb is imported from telemetry_v1
/*
func (s *TelemetryServer) TelemetryStream(stream pb.TelemetryService_TelemetryStreamServer) error {
    for {
        req, err := stream.Recv()
        if err == io.EOF {
            return nil
        }
        if err != nil {
            return err
        }
        
        // Logic
        if req.Speed > 180 {
            stream.Send(&pb.Alert{Message: "Too fast!"})
        }
    }
}
*/
