package grpc_handler

import (
	"bytes"
	"context"

	carv1 "github.com/course/tasks/api/car/v1"
	"github.com/xuri/excelize/v2"
)

type MaintenanceHandler struct {
	carv1.UnimplementedMaintenanceServiceServer
	// repo *db.Queries
}

// Task 10: Excel export
func (h *MaintenanceHandler) ExportWorkOrders(ctx context.Context, req *carv1.ExportWorkOrdersRequest) (*carv1.ExportWorkOrdersResponse, error) {
	// 1. Fetch data from DB
	// orders := h.repo.GetWorkOrders(ctx)

	// 2. Generate Excel file
	f := excelize.NewFile()
	defer f.Close()

	sheet := "Sheet1"
	f.SetCellValue(sheet, "A1", "Work Order ID")
	f.SetCellValue(sheet, "B1", "Car VIN")
	f.SetCellValue(sheet, "C1", "Status")

	// for i, order := range orders {
	//     row := i + 2
	//     f.SetCellValue(sheet, fmt.Sprintf("A%d", row), order.ID)
	//     f.SetCellValue(sheet, fmt.Sprintf("B%d", row), order.CarVin)
	//     f.SetCellValue(sheet, fmt.Sprintf("C%d", row), order.Status)
	// }

	// Dummy data
	f.SetCellValue(sheet, "A2", 1)
	f.SetCellValue(sheet, "B2", "VIN123456")
	f.SetCellValue(sheet, "C2", "COMPLETED")

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}

	return &carv1.ExportWorkOrdersResponse{
		FileUrl:     "generated", // In a real app we might save it to S3 and return a URL
		FileContent: buf.Bytes(),
	}, nil
}
