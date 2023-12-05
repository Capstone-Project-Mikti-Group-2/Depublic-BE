BEGIN;

CREATE TABLE IF NOT EXISTS "public"."topup" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "user_id" SERIAL NOT NULL,
    "nominal" BIGINT,
    "snap_url" varchar(255),
    "status" INT NOT NULL,
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
);

COMMIT;