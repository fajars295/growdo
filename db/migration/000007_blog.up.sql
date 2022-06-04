CREATE TABLE "blog" (
  "id" bigserial PRIMARY KEY,
  "title" varchar(225) NOT NULL,
  "images" varchar(225) NOT NULL,
  "tag" varchar(255),
  "deskripsi" TEXT NOT NULL,
  "status" BOOLEAN NOT NULL DEFAULT (false),
  "created_by" bigint NOT NULL,
  "master_data_id" INT CONSTRAINT blog_fk_master_data_id REFERENCES master_data (id) ON UPDATE CASCADE ON DELETE CASCADE,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);