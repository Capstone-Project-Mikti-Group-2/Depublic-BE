BEGIN;

CREATE TABLE IF NOT EXISTS "public"."tickets" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "user_id" INT NOT NULL,
    "event_id" INT NOT NULL,
    "transaction_id" INT NOT NULL,
    "quantity" INT NOT NULL,
    "total" INT NOT NULL,
    "is_paid" BOOLEAN NOT NULL DEFAULT false,
    "book_by" VARCHAR(255) NOT NULL,
    "update_by" VARCHAR(255) NOT NULL,
    "delete_by" VARCHAR(255),
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    "deleted_at" timestamptz (6),
    FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("event_id") REFERENCES "public"."events" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("transaction_id") REFERENCES "public"."transactions" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);

COMMIT;