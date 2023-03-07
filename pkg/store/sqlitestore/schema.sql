PRAGMA foreign_keys = ON;
PRAGMA defer_foreign_keys = FALSE;

CREATE TABLE IF NOT EXISTS users (
	username   TEXT PRIMARY KEY NOT NULL,
	password   TEXT NOT NULL,
	email      TEXT NOT NULL,
	phone      TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS projects (
	name    TEXT PRIMARY KEY NOT NULL,
	descr   TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS work (
	id       INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	project  TEXT NOT NULL,
	ts       INTEGER NOT NULL,
	duration INTEGER NOT NULL,
	descr    TEXT NOT NULL,

	FOREIGN KEY(username) REFERENCES users(username),
	FOREIGN KEY(project) REFERENCES projects(name)
);
CREATE INDEX work_project ON work(project);
CREATE INDEX work_username ON work(username);
CREATE INDEX work_ts ON work(ts);

CREATE TABLE IF NOT EXISTS snippets (
	id       INTEGER PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	ts       INTEGER NOT NULL,
	contents TEXT NOT NULL,

	FOREIGN KEY(username) REFERENCES users(username)
);
CREATE INDEX snippets_username ON snippets(username);
CREATE INDEX snippets_ts ON snippets(ts);

