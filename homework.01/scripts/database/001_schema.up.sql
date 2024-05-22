-- Create default schema
CREATE SCHEMA IF NOT EXISTS public;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp"
WITH
  SCHEMA public;

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';

CREATE TABLE IF NOT EXISTS "users" (
  "id" uuid PRIMARY KEY DEFAULT (public.uuid_generate_v4()),
  "hash" TEXT NOT NULL,
  "first_name" TEXT DEFAULT null,
  "second_name" TEXT DEFAULT null,
  "birthdate" TEXT DEFAULT null,
  "biography" TEXT DEFAULT null,
  "city" TEXT DEFAULT null
);
