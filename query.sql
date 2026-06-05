-- name: CreateUser :one
INSERT INTO users (name) VALUES ($1) RETURNING *;

-- name: AddDocument :one
INSERT INTO documents (user_id, title, content) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUserDocuments :many
SELECT * FROM documents WHERE user_id = $1;

-- name: AssembleCar :exec
UPDATE cars SET engine_id = $1, transmission_id = $2 WHERE vin = $3;

-- name: GetCarSpec :one
SELECT c.vin, c.brand, c.year, e.model as engine_model, e.horsepower as engine_hp, t.type as transmission_type 
FROM cars c
LEFT JOIN engines e ON c.engine_id = e.id
LEFT JOIN transmissions t ON c.transmission_id = t.id
WHERE c.vin = $1;

-- name: GetOwnerDashboard :many
SELECT o.id as owner_id, o.name as owner_name, c.vin, c.brand, sr.date as last_service_date, sr.description
FROM owners o
JOIN cars c ON o.id = c.owner_id
LEFT JOIN service_records sr ON c.vin = sr.car_vin
WHERE o.id = $1
ORDER BY sr.date DESC;

-- name: CreateBooking :one
INSERT INTO bookings (car_vin, start_date, end_date) VALUES ($1, $2, $3) RETURNING *;

-- name: CheckCarAvailability :one
-- Task 8: Check for overlapping bookings with FOR UPDATE to lock
SELECT COUNT(*) FROM bookings 
WHERE car_vin = $1 AND start_date <= $3 AND end_date >= $2
FOR UPDATE;

-- name: LogChange :exec
INSERT INTO changelogs (entity_type, entity_id, action) VALUES ($1, $2, $3);

-- name: GetWorkOrders :many
SELECT * FROM work_orders;
