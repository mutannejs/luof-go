CREATE TABLE link (
    uid_link CHARACTER(16) PRIMARY KEY,
    url TEXT,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    use_markdown BOOLEAN,
    created_at DATETIME NOT NULL,
    update_at DATETIME NOT NULL
);
