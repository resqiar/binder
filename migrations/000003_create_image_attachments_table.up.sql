CREATE TABLE IF NOT EXISTS image_attachments (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    URL   TEXT UNIQUE NOT NULL,
    extension_id UUID,

    FOREIGN KEY(extension_id) REFERENCES extensions(id)
);

