   BEGIN;

   -- Create users table
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       email VARCHAR(100) NOT NULL UNIQUE,
       password VARCHAR(255) NOT NULL,
       role VARCHAR(10) CHECK (role IN ('Admin', 'Sewa', 'Penyewa')) NOT NULL,
       balance FLOAT DEFAULT 0.0,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       deleted_at TIMESTAMP NULL
   );

   -- Create motor table
   CREATE TABLE motor (
       id SERIAL PRIMARY KEY,
       owner_id INT NOT NULL,
       motor_name VARCHAR(100) NOT NULL,
       license_plate VARCHAR(20) NOT NULL UNIQUE,
       rental_price_per_day FLOAT NOT NULL,
       status VARCHAR(10) CHECK (status IN ('Available', 'Rented')) DEFAULT 'Available' NOT NULL,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       FOREIGN KEY (owner_id) REFERENCES users(id)
   );

   -- Create transactions table
   CREATE TABLE transactions (
       id SERIAL PRIMARY KEY,
       renter_id INT NOT NULL,
       motor_id INT NOT NULL,
       start_date DATE NOT NULL,
       end_date DATE NOT NULL,
       total_cost FLOAT NOT NULL,
       status VARCHAR(10) CHECK (status IN ('Successful', 'Failed')) DEFAULT 'Successful' NOT NULL,
       admin_fee FLOAT NOT NULL,
       FOREIGN KEY (renter_id) REFERENCES users(id),
       FOREIGN KEY (motor_id) REFERENCES motor(id)
   );

   -- Create topup table
   CREATE TABLE topup (
       id VARCHAR(255) PRIMARY KEY,
       user_id INT NOT NULL,
       amount INT NOT NULL,
       status INT NOT NULL,
       snap_url VARCHAR(255),
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       FOREIGN KEY (user_id) REFERENCES users(id)
   );

   COMMIT;