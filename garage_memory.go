package repository

import (
    "context"
    "errors"
    "sync"
    "github.com/course/tasks/internal/domain"
)

// Task 1: In-memory map
type GarageMemoryRepo struct {
    mu   sync.RWMutex
    cars map[string]domain.Car
}

func NewGarageMemoryRepo() *GarageMemoryRepo {
    return &GarageMemoryRepo{
        cars: make(map[string]domain.Car),
    }
}

func (r *GarageMemoryRepo) CreateCar(ctx context.Context, car domain.Car) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    if _, exists := r.cars[car.VIN]; exists {
        return errors.New("car already exists")
    }
    r.cars[car.VIN] = car
    return nil
}

func (r *GarageMemoryRepo) GetCar(ctx context.Context, vin string) (domain.Car, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    car, exists := r.cars[vin]
    if !exists {
        return domain.Car{}, errors.New("car not found")
    }
    return car, nil
}
