package storage

import (
	"database/sql"
	"fmt"
)

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

func RenameMark(id int, name string) bool {
	flag := false
	dbfile := GetMarkTable(true)

	Connection(dbfile, func(db *sql.DB) {
		if _, err := db.Exec(RENAME_MARK_SQL, name, id); err != nil {
			fmt.Println("err:\n", err)
		} else {
			flag = true
		}
	})

	return flag
}

func QueryMark(id int) (Mark, error) {
	dbfile := GetMarkTable(true)
	var mark Mark
	var err error = nil

	Connection(dbfile, func(db *sql.DB) {
		rows, e := db.Query(QUERY_MARK_SQL, id)
		if e != nil {
			fmt.Println("err:\n", e)
			err = e
			return
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&mark.Id, &mark.Mark, &mark.Description, &mark.Color, &mark.Icon, &mark.Sort, &mark.CreatedAt, &mark.ModifyAt)
			if err != nil {
				fmt.Println("err:\n", err)
			}
			break
		}
	})

	return mark, err
}

func ChangeMark(id int, mark *Mark) bool {
	flag := false
	dbfile := GetMarkTable(true)

	Connection(dbfile, func(db *sql.DB) {
		if _, err := db.Exec(CHANGE_MARK_SQL, mark.Mark, mark.Description, mark.Color, mark.Icon, id); err != nil {
			fmt.Println("err:\n", err)
		} else {
			flag = true
		}
	})

	return flag
}
