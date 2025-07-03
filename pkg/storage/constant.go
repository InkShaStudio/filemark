package storage

const TABLE_FILE_NAME = "mark.db"
const CREATE_MARK_TABLE_SQL = `
	CREATE TABLE IF NOT EXISTS mark (
		id INTEGER PRIMARY KEY,
		mark TEXT,
		description TEXT,
		color TEXT,
		icon TEXT,
		sort INTEGER DEFAULT 1,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		modify_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
`

const CREATE_FILE_TABLE_SQL = `
	CREATE TABLE IF NOT EXISTS files (
		id INTEGER PRIMARY KEY,
		file_path TEXT,
		dir TEXT,
		sha256 TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		modify_at DATETIME DEFAULT CURRENT_TIMESTAMP,

		UNIQUE(file_path)
	);
`

const CREATE_RELATION_TABLE_SQL = `
	CREATE TABLE IF NOT EXISTS file_marks (
		id INTEGER PRIMARY KEY,
		file_id INTEGER,
		mark_id INTEGER,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		modify_at DATETIME DEFAULT CURRENT_TIMESTAMP,

		FOREIGN KEY (file_id) REFERENCES files(id) ON DELETE CASCADE,
		FOREIGN KEY (mark_id) REFERENCES mark(id) ON DELETE CASCADE,
		UNIQUE(file_id, mark_id)
	);
`

const INSERT_MARK_SQL = `INSERT INTO mark (mark, description, color, icon) VALUES(?,?,?,?);`
const DELETE_MARK_SQL = `DELETE FROM mark WHERE id = ?;`
const RENAME_MARK_SQL = `UPDATE mark SET mark=? WHERE id=?;`
const QUERY_MARK_SQL = `SELECT * FROM mark WHERE id=?;`
const CHANGE_MARK_SQL = `UPDATE mark SET mark=?, description=?, color=?, icon=? WHERE id=?;`

const INSERT_FILE_SQL = `INSERT INTO files (file_path, dir, sha256) VALUES(?,?,?);`

const QUERY_ALL_FILES_BY_DIR = `SELECT * FROM files WHERE dir = ?;`
const QUERY_FILE_BY_SHA256 = `SELECT * FROM files WHERE sha256 = ?;`
const QUERY_FILE_BY_PATH = `SELECT * FROM files WHERE file_path = ?;`

const CHANGE_FILE_PATH = `UPDATE files SET file_path = ? WHERE id = ?;`
const CHANGE_FILE_DIR = `UPDATE files SET dir = ? WHERE id = ?;`
const CHANGE_FILE_SHA256 = `UPDATE files SET sha256 = ? WHERE id = ?;`

const QUERY_FILE_ALL_MARKS_BY_ID_SQL = `SELECT m.* FROM mark m JOIN file_marks fm ON m.id = fm.mark_id WHERE fm.file_id = ?;`
const QUERY_FILE_ALL_MARKS_BY_PATH_SQL = `SELECT m.* FROM mark m JOIN file_marks fm ON m.id = fm.mark_id JOIN files f ON fm.file_id = f.id WHERE f.file_path = ?;`

const QUERY_MARK_ALL_FILES_SQL = `SELECT f.* FROM files f JOIN file_marks fm ON f.id = fm.file_id WHERE fm.mark_id = ?;`

const INSERT_FILE_MARK_SQL = `INSERT INTO file_marks (file_id, mark_id) VALUES(?,?);`
