CREATE TABLE public.music
(
    uuid         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at   TIMESTAMPTZ      DEFAULT NOW(),
    history_uuid UUID REFERENCES public.history (uuid) ON DELETE CASCADE ON UPDATE CASCADE,
    music_name   VARCHAR,
    path         VARCHAR,
    thumbnail    VARCHAR,
    artist       VARCHAR
);