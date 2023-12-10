BEGIN;

CREATE TABLE IF NOT EXISTS "public"."transactions" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "user_id" INT NOT NULL,
    "order_id" VARCHAR(255) NOT NULL,
    "amount" INT NULL,
    "status" varchar(255) NOT NULL,
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    "deleted_at" timestamptz (6),
    FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);


COMMIT;