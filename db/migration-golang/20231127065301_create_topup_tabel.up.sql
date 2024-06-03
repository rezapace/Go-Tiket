BEGIN;
CREATE TABLE IF NOT EXISTS "public"."topup" (
    "id" varchar(255) NOT NULL PRIMARY KEY,
    "user_id" INT,
    "amount" INT,
    "status" INT,
    "snap_url" varchar(255),
    "created_at" TIMESTAMP,
    "updated_at" TIMESTAMP
);
COMMIT;