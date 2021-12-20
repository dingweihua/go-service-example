CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    username text,
    mobile text,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    disabled_at timestamptz
    );
