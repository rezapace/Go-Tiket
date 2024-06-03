BEGIN;
CREATE TABLE IF NOT EXISTS "public"."tickets" (
    ID SERIAL NOT NULL PRIMARY KEY,
    Image TEXT,
    Location TEXT,
    Date DATE,
    Title TEXT,
    Description TEXT,
    Price INT,
    Status TEXT DEFAULT 'available',
    Quota INT,
    Category TEXT,
    Terjual INT,
    Created_At TIMESTAMP,
    Updated_At TIMESTAMP,
    Deleted_At TIMESTAMP
);
COMMIT;