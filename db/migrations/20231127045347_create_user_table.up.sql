BEGIN;


CREATE TABLE IF NOT EXISTS "public"."users" (
    "id" SERIAL NOT NULL PRIMARY KEY,
    "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "email" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
    "number" varchar(255),
    "role" varchar(255),
    "saldo" INTEGER NOT NULL,
    "created_at" timestamptz (6) NOT NULL,
    "updated_at" timestamptz (6) NOT NULL,
    "deleted_at" timestamptz (6)
);


COMMIT;