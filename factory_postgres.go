package repository

import (
    "context"
    "database/sql"
)

// Task 3: AssembleCar with DB Transactions
type FactoryRepo struct {
    db *sql.DB
}

func NewFactoryRepo(db *sql.DB) *FactoryRepo {
    return &FactoryRepo{db: db}
}

func (r *FactoryRepo) AssembleCar(ctx context.Context, carID, engineID, transID int) error {
    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback() // Will be ignored if tx.Commit() is called

    _, err = tx.ExecContext(ctx, "UPDATE cars SET engine_id = $1 WHERE id = $2", engineID, carID)
    if err != nil {
        return err
    }

    _, err = tx.ExecContext(ctx, "UPDATE cars SET trans_id = $1 WHERE id = $2", transID, carID)
    if err != nil {
        return err
    }

    return tx.Commit()
}
