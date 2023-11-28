BEGIN;

CREATE TABLE IF NOT EXISTS "public"."profiles" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "user_id" SERIAL NOT NULL,
    "address" varchar(255),
    "avatar" bytea,
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    "deleted_at" timestamptz (6),
    FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id")
);

COMMIT;