package repository

import (
    "context"
    "database/sql"
)

// Task 10: DB Hooks / Changelog
type WorkOrderRepo struct {
    db *sql.DB
}

func (r *WorkOrderRepo) UpdateStatus(ctx context.Context, orderID int, newStatus string) error {
    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    // 1. Update the order
    _, err = tx.ExecContext(ctx, "UPDATE work_orders SET status = $1 WHERE id = $2", newStatus, orderID)
    if err != nil {
        return err
    }

    // 2. Add Changelog (Hook behavior)
    _, err = tx.ExecContext(ctx, "INSERT INTO changelog (entity_type, entity_id, action, timestamp) VALUES ('WORK_ORDER', $1, 'STATUS_UPDATE', NOW())", orderID)
    if err != nil {
        return err
    }

    return tx.Commit()
}
