BEGIN;

CREATE TABLE IF NOT EXISTS "public". "orders" (
    ID SERIAL PRIMARY KEY,
    ticket_id INT,
    user_id INT,
    quantity INT,
    total INT,
    status VARCHAR(255),
    order_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    order_by VARCHAR(255),
    update_by VARCHAR(255),
    delete_by VARCHAR(255),
    FOREIGN KEY (ticket_id) REFERENCES tickets(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);

COMMIT;

