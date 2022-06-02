CREATE TABLE "banner" (
    "id" bigserial PRIMARY KEY,
    "images" varchar NOT NULL,
    "status" BOOLEAN NOT NULL DEFAULT (false),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
);