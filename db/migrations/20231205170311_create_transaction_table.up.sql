BEGIN;

CREATE TABLE IF NOT EXISTS "public"."transactions" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "user_id" INT NOT NULL,
    "order_id" VARCHAR(255) NOT NULL,
    "status" VARCHAR(255) NOT NULL,
    "amount" INT NULL,
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    "deleted_at" timestamptz (6)
);

COMMIT;