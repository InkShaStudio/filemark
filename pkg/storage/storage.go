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
const DELETE_MARK_SQL = `DELETE FROM mark WHERE id = ?;`

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
			fmt.Println("err:\n", err)
		}
	})
}

func QueryMarks() []Mark {
	marks := make([]Mark, 0)
	dbfile := GetMarkTable(true)

	Connection(dbfile, func(db *sql.DB) {
		rows, err := db.Query("SELECT * FROM mark")
		if err != nil {
			fmt.Println("err:\n", err)
		}
		defer rows.Close()
		for rows.Next() {
			var mark Mark
			err = rows.Scan(&mark.Id, &mark.Mark, &mark.Description, &mark.Color, &mark.Icon, &mark.Sort, &mark.CreatedAt, &mark.ModifyAt)
			if err != nil {
				fmt.Println("err:\n", err)
			}
			marks = append(marks, mark)
		}
	})

	return marks
}

func InsertMark(data CreateMark) bool {
	flag := false
	dbfile := GetMarkTable(true)

	Connection(dbfile, func(db *sql.DB) {
		result, err := db.Exec(INSERT_MARK_SQL, data.Mark, data.Description, data.Color, data.Icon)
		if err != nil {
			fmt.Println("err:\n", err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			fmt.Println("err:\n", err)
		}
		if id > 0 {
			flag = true
		}
	})

	return flag
}

func RemoveMark(id int) bool {
	flag := false
	dbfile := GetMarkTable(true)

	Connection(dbfile, func(db *sql.DB) {
		_, err := db.Exec(DELETE_MARK_SQL, id)
		if err != nil {
			fmt.Println("err:\n", err)
		} else {
			flag = true
		}
	})

	return flag
}
