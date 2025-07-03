package storage

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Connection(dbfile string, callback func(db *sql.DB)) {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()
	callback(db)
}

func GetMarkTable(root bool) string {
	if root {
		home, err := os.UserHomeDir()

		if err != nil {
			panic("No Get Home Dir!")
		}

		return home + "/" + TABLE_FILE_NAME
	}

	return TABLE_FILE_NAME
}

func CreateTable() {
	dbfile := GetMarkTable(true)

	Connection(dbfile, func(db *sql.DB) {
		_, err := db.Exec(CREATE_MARK_TABLE_SQL)
		if err != nil {
			fmt.Println("create mrak table err:\n", err)
		}
		_, err = db.Exec(CREATE_FILE_TABLE_SQL)
		if err != nil {
			fmt.Println("create file table err:\n", err)
		}

		_, err = db.Exec(CREATE_RELATION_TABLE_SQL)
		if err != nil {
			fmt.Println("create relation table err:\n", err)
		}
	})
}
