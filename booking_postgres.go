package repository

import (
    "context"
    "database/sql"
    "errors"
    "time"
)

// Task 8: Optimistic/Pessimistic locking
type BookingRepo struct {
    db *sql.DB
}

func (r *BookingRepo) CreateBooking(ctx context.Context, carID int, start, end time.Time) error {
    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    // Pessimistic Lock on the car
    var lockedCarID int
    err = tx.QueryRowContext(ctx, "SELECT id FROM cars WHERE id = $1 FOR UPDATE", carID).Scan(&lockedCarID)
    if err != nil {
        return err
    }

    // Check overlaps
    var count int
    err = tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM bookings WHERE car_id = $1 AND start_time < $2 AND end_time > $3", carID, end, start).Scan(&count)
    if err != nil || count > 0 {
        return errors.New("car is already booked for these dates")
    }

    _, err = tx.ExecContext(ctx, "INSERT INTO bookings (car_id, start_time, end_time) VALUES ($1, $2, $3)", carID, start, end)
    if err != nil {
        return err
    }

    return tx.Commit()
}
