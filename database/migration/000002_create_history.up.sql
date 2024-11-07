CREATE TABLE public.history
(
    uuid          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at    TIMESTAMPTZ      DEFAULT NOW(),
    user_uuid     UUID REFERENCES public.users (uuid) ON DELETE CASCADE ON UPDATE CASCADE,
    mood          VARCHAR,
    playlist_name VARCHAR,
    path          VARCHAR
);