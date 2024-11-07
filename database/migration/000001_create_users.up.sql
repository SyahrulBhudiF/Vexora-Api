CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE public.users
(
    uuid            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at      TIMESTAMPTZ      DEFAULT NOW(),
    username        VARCHAR        NOT NULL,
    name            VARCHAR,
    email           VARCHAR UNIQUE NOT NULL,
    password        VARCHAR        NOT NULL,
    profile_picture VARCHAR
);