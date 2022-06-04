CREATE TABLE "master_data" (
    "id" bigserial PRIMARY KEY,
    "type" VARCHAR NOT NULL,
    "value" VARCHAR NOT NULL,
    "status" BOOLEAN NOT NULL DEFAULT (false),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);