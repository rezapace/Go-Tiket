BEGIN;

-- Insert sample data into users table
INSERT INTO users (name, email, number, password, role, balance, created_at, updated_at) VALUES 
('Admin User', 'admin@example.com', '1234567890', 'adminpassword', 'Admin', 0.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Sewa User', 'sewa@example.com', '0987654321', 'sewapassword', 'Sewa', 100.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('Penyewa User', 'penyewa@example.com', '1122334455', 'penyewapassword', 'Penyewa', 50.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert sample data into motor table
INSERT INTO motor (owner_id, motor_name, license_plate, rental_price_per_day, status, created_at, updated_at) VALUES 
(2, 'Yamaha NMAX', 'B1234XYZ', 150.0, 'Available', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
(2, 'Honda Vario', 'B5678ABC', 100.0, 'Rented', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

-- Insert sample data into transactions table
INSERT INTO transactions (renter_id, motor_id, start_date, end_date, total_cost, status, admin_fee) VALUES 
(3, 1, '2023-12-01', '2023-12-05', 600.0, 'Successful', 50.0),
(3, 2, '2023-12-10', '2023-12-15', 500.0, 'Failed', 50.0);

-- Insert sample data into topup table
INSERT INTO topup (id, user_id, amount, status, snap_url, created_at, updated_at) VALUES 
('topup1', 2, 200, 1, 'http://example.com/snap1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
('topup2', 3, 300, 1, 'http://example.com/snap2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

COMMIT;