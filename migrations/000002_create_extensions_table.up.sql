CREATE TABLE IF NOT EXISTS extensions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    title   TEXT NOT NULL,
    description   TEXT,
    youtube_url TEXT,
    author_id UUID,

    FOREIGN KEY(author_id) REFERENCES users(id)
);
