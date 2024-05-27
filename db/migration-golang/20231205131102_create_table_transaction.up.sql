BEGIN;

CREATE TABLE IF NOT EXISTS "public"."transactions" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "order_id" varchar(255) NOT NULL,
    "user_id" INT NOT NULL,
    "amount" INT NULL,
    "status" varchar(255) NOT NULL
    -- "id" varchar(255) NOT NULL PRIMARY KEY,
    -- "user_id" INT,
    -- "amount" INT,
    -- "status" INT,
    -- "snap_url" varchar(255),
    -- "created_at" TIMESTAMP,
    -- "updated_at" TIMESTAMP
);

COMMIT;