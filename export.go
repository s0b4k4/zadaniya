package rest

import (
    "net/http"
    "strconv"
    // "github.com/xuri/excelize/v2"
)

// Task 10: Export WorkOrders to Excel
func ExportWorkOrders(w http.ResponseWriter, r *http.Request) {
    /*
    f := excelize.NewFile()
    defer f.Close()

    f.SetCellValue("Sheet1", "A1", "Order ID")
    f.SetCellValue("Sheet1", "B1", "Status")

    // Fetch orders from DB and loop
    for i, order := range orders {
        row := strconv.Itoa(i + 2)
        f.SetCellValue("Sheet1", "A"+row, order.ID)
        f.SetCellValue("Sheet1", "B"+row, order.Status)
    }

    w.Header().Set("Content-Type", "application/octet-stream")
    w.Header().Set("Content-Disposition", "attachment; filename=work_orders.xlsx")
    f.Write(w)
    */
}
