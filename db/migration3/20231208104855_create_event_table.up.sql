BEGIN;

CREATE TABLE IF NOT EXISTS "public"."events" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "description" TEXT,
    "location" VARCHAR(255),
    "image" BYTEA,
    "price" INT NOT NULL,
    "quantity" INT NOT NULL,
    "available" BOOLEAN NOT NULL,
    "start_date" DATE NOT NULL,
    "end_date" DATE NOT NULL,
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    "deleted_at" timestamptz (6)
);

COMMIT;