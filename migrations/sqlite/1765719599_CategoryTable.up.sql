CREATE TABLE category (
	uid_category CHARACTER(16) PRIMARY KEY,
	name VARCHAR(200) NOT NULL,
	description TEXT,
	use_markdown BOOLEAN,
	uid_father CHARACTER(16),
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL,
	FOREIGN KEY (uid_father) REFERENCES category(uid_category)
);
