CREATE TABLE category (
	uid_category CHARACTER(16) PRIMARY KEY,
	name VARCHAR(200) NOT NULL,
	description TEXT,
	use_markdown BOOLEAN,
	created_at DATETIME NOT NULL,
	updated_at DATETIME NOT NULL
);
