CREATE TABLE IF NOT EXISTS extensions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    slug TEXT UNIQUE NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    title   TEXT NOT NULL,
    description   TEXT,
    youtube_url TEXT,
    code   TEXT,
    author_id UUID,
    
    search_vector TSVECTOR GENERATED ALWAYS AS (
        to_tsvector('english', slug || ' ' || title || ' ' || COALESCE(description, ' '))
    ) STORED,

    FOREIGN KEY(author_id) REFERENCES users(id)
);

CREATE INDEX IF NOT EXISTS extensions_search_vector ON extensions USING GIN (search_vector);
