BEGIN;

CREATE TABLE IF NOT EXISTS "public"."transaction" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "user_id" NOT NULL,
    "event_id" NULL,
    "topup_id"  NULL,
    "amount" DECIMAL NOT NULL,
    "transaction_type" varchar(255) NOT NULL,
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("event_id") REFERENCES "public"."events" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY ("topup_id") REFERENCES "public"."topup" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);

COMMIT;