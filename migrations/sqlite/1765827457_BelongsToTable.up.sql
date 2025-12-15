CREATE TABLE belongs_to (
    uid_link CHARACTER(16) NOT NULL,
    uid_category CHARACTER(16) NOT NULL,
    is_main BOOLEAN,
    inserted_at DATETIME NOT NULL,
    FOREIGN KEY (uid_link) REFERENCES link(uid_link),
    FOREIGN KEY (uid_category) REFERENCES category(uid_category),
    PRIMARY KEY (uid_link, uid_category)
);
