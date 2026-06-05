-- Task 2: Users and Documents
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE documents (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    title VARCHAR(255) NOT NULL,
    content TEXT
);

-- Task 3, 6, 8, 10: Car, Engine, Transmission, Owner, ServiceRecord, Part, WorkOrder, Booking
CREATE TABLE cars (
    vin VARCHAR(17) PRIMARY KEY,
    brand VARCHAR(100) NOT NULL,
    year INT NOT NULL,
    engine_id INT,
    transmission_id INT,
    owner_id INT
);

CREATE TABLE engines (
    id SERIAL PRIMARY KEY,
    model VARCHAR(100) NOT NULL,
    horsepower INT NOT NULL
);

CREATE TABLE transmissions (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) NOT NULL
);

CREATE TABLE owners (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE service_records (
    id SERIAL PRIMARY KEY,
    car_vin VARCHAR(17) NOT NULL REFERENCES cars(vin),
    date DATE NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE bookings (
    id SERIAL PRIMARY KEY,
    car_vin VARCHAR(17) NOT NULL REFERENCES cars(vin),
    start_date DATE NOT NULL,
    end_date DATE NOT NULL
);

CREATE TABLE parts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

CREATE TABLE work_orders (
    id SERIAL PRIMARY KEY,
    car_vin VARCHAR(17) NOT NULL REFERENCES cars(vin),
    date DATE NOT NULL,
    status VARCHAR(50) NOT NULL
);

CREATE TABLE changelogs (
    id SERIAL PRIMARY KEY,
    entity_type VARCHAR(50) NOT NULL,
    entity_id VARCHAR(50) NOT NULL,
    action VARCHAR(50) NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT NOW()
);
