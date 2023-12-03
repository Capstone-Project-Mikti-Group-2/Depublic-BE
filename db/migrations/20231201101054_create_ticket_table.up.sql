BEGIN;

CREATE TABLE IF NOT EXISTS "public"."tickets" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "user_id" SERIAL NOT NULL,
    "event_id" SERIAL NOT NULL,
    "quantity" INT NOT NULL,
    "total" INT NOT NULL,
    "isPaid" BOOLEAN NOT NULL,
    "book_by" VARCHAR(255) NOT NULL,
    "updated_by" VARCHAR(255) NOT NULL,
    "deleted_by" VARCHAR(255),
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    "deleted_at" timestamptz (6),
    FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("event_id") REFERENCES "public"."events" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);

COMMIT;