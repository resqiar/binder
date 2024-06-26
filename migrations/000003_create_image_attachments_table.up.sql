CREATE TABLE IF NOT EXISTS image_attachments (
    id TEXT UNIQUE PRIMARY KEY NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    URL   TEXT NOT NULL,
    extension_id UUID,

    FOREIGN KEY(extension_id) REFERENCES extensions(id)
);
