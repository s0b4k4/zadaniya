package main

import (
    "context"
    "fmt"
    "sync"
    "errors"
    pb "github.com/yourusername/garage/api" // Подключаем сгенерированный код
)

// server - это наша структура, которая будет выполнять работу
type server struct {
    pb.UnimplementedGarageServiceServer // Магическая строчка для совместимости
    
    mu   sync.RWMutex // Замочек, чтобы два пользователя не подрались за данные
    cars map[string]*pb.Car // Наша "база данных" в памяти (словарь: VIN -> Автомобиль)
}

func newServer() *server {
    return &server{
        cars: make(map[string]*pb.Car), // Создаем пустой словарь при старте
    }
}
