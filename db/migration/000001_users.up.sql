CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "email" VARCHAR NOT NULL UNIQUE,
    "password" text NOT NULL,
    "date_of_birth" DATE NOT NULL,
    "phone" VARCHAR NOT NULL,
    "role" BIGINT DEFAULT (1),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);