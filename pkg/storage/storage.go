package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

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
const INSERT_MARK_SQL = `INSERT INTO mark (mark, description, color, icon) VALUES(?,?,?,?);`

func Connection(dbfile string, callback func(db *sql.DB)) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()
	callback(db)
}

func CreateTable() {
	home, err := os.UserHomeDir()

	if err != nil {
		panic("No Get Home Dir!")
	}

	dbfile := home + "/" + TABLE_FILE_NAME

	Connection(dbfile, func(db *sql.DB) {
		_, err = db.Exec(CREATE_MARK_TABLE_SQL)
		if err != nil {
			fmt.Println("err:\n", err)
		}
	})
}
