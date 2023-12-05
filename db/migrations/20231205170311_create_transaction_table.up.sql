BEGIN;

CREATE TABLE IF NOT EXISTS "public"."transaction" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "event_id" SERIAL NOT NULL,
    "user_id" INT NOT NULL,
    "order_id" VARCHAR(255) NOT NULL,
    "amount" INT NULL,
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    "deleted_at" timestamptz (6),
    FOREIGN KEY ("event_id") REFERENCES "public"."event" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
    FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);

COMMIT;