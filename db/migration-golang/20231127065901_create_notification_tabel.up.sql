BEGIN;
CREATE TABLE IF NOT EXISTS "public"."notifications" (
    ID SERIAL NOT NULL PRIMARY KEY,
    Type TEXT,
    Message TEXT,
    Is_Read BOOLEAN DEFAULT FALSE,
    Created_At TIMESTAMP,
    Updated_At TIMESTAMP,
    Deleted_At TIMESTAMP
);
COMMIT;