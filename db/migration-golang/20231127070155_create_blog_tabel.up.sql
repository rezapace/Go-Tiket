BEGIN;
CREATE TABLE IF NOT EXISTS "public". "blogs" (
    ID SERIAL PRIMARY KEY,
    Image TEXT,
    Date DATE,
    Title TEXT,
    Description TEXT,
    Created_At TIMESTAMP,
    Updated_At TIMESTAMP,
    Deleted_At TIMESTAMP
);
COMMIT;