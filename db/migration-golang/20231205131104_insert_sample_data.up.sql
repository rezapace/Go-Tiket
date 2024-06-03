   BEGIN;

   INSERT INTO users (name, email, password, role, balance, created_at, updated_at) VALUES 
   ('Admin User', 'admin@example.com', 'adminpassword', 'Admin', 0.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
   ('Sewa User', 'sewa@example.com', 'sewapassword', 'Sewa', 100.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
   ('Penyewa User', 'penyewa@example.com', 'penyewapassword', 'Penyewa', 50.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

   INSERT INTO motor (owner_id, motor_name, license_plate, rental_price_per_day, status, created_at, updated_at) VALUES
   (2, 'Yamaha NMax', 'AB123CD', 100000, 'Available', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP),
   (2, 'Honda Vario', 'AB456EF', 80000, 'Available', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

   INSERT INTO transactions (renter_id, motor_id, start_date, end_date, total_cost, status, admin_fee) VALUES
   (3, 1, '2024-06-01', '2024-06-03', 200000, 'Successful', 20000);

   INSERT INTO topup (id, user_id, amount, status, snap_url, created_at, updated_at) VALUES
   ('1', 3, 100000, 1, 'http://example.com/snap/1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

   COMMIT;