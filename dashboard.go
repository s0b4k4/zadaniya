package rest

import (
    "context"
    "database/sql"
    "encoding/json"
    "net/http"
)

// Task 6: Dashboard Response Struct
type DashboardResponse struct {
    OwnerName       string `json:"owner_name"`
    CarBrand        string `json:"car_brand"`
    LastServiceDate string `json:"last_service_date"`
}

type DashboardHandler struct {
    db *sql.DB
}

// REST Dashboard endpoint (Complex Query for frontend)
func (h *DashboardHandler) GetDashboard(w http.ResponseWriter, r *http.Request) {
    ownerID := r.URL.Query().Get("owner_id")
    
    query := `
        SELECT o.name, c.brand, s.last_service_date
        FROM owners o
        LEFT JOIN cars c ON c.owner_id = o.id
        LEFT JOIN service_records s ON s.car_id = c.id
        WHERE o.id = $1
    `
    
    rows, err := h.db.QueryContext(context.Background(), query, ownerID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var dashboard []DashboardResponse
    for rows.Next() {
        var dr DashboardResponse
        if err := rows.Scan(&dr.OwnerName, &dr.CarBrand, &dr.LastServiceDate); err != nil {
            continue
        }
        dashboard = append(dashboard, dr)
    }

    json.NewEncoder(w).Encode(dashboard)
}
