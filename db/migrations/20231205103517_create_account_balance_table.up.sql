BEGIN;

CREATE TABLE IF NOT EXISTS "public"."account_balance" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "user_id" SERIAL NOT NULL,
    "event_id" SERIAL NOT NULL,
    "balance" DECIMAL NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
);


COMMIT;