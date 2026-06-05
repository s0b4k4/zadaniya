package domain

type Car struct {
    VIN   string
    Brand string
    Year  int
}

type Engine struct {
    ID         int
    Model      string
    Horsepower int
}

type Transmission struct {
    ID   int
    Type string
}
