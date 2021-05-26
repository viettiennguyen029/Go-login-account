CREATE TABLE "users" (
  "address" varchar PRIMARY KEY NOT NULL,
  "image_url" varchar NOT NULL,
  "created_time" timestamp NOT NULL,
  "updated_time" timestamp NOT NULL
);

CREATE TABLE "tokens" (
  "token" varchar PRIMARY KEY NOT NULL
);