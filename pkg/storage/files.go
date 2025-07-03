package storage

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func QueryFiles() []FileMark {
	result := make([]FileMark, 0)
	dbfile := GetMarkTable(true)

	Connection(dbfile, func(db *sql.DB) {
		rows, err := db.Query("select * from files")
		if err != nil {
			fmt.Println("query files err:\n", err)
		}
		defer rows.Close()
		for rows.Next() {
			var file FileMark
			err := rows.Scan(&file.ID, &file.FilePath, &file.Dir, &file.Sha256, &file.CreatedAt, &file.ModifyAt)
			if err != nil {
				fmt.Println("scan files err:\n", err)
			}
			result = append(result, file)
		}
	})

	return result
}

func CalculateFileSHA256(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func InsertFile(path string) bool {
	flag := false

	dbfile := GetMarkTable(true)

	Connection(dbfile, func(db *sql.DB) {

		if hash, err := CalculateFileSHA256(path); err != nil {
			fmt.Println("calculate file sha256 failure\n", err)
		} else {
			if _, err := db.Exec(INSERT_FILE_SQL, path, filepath.Dir(path), hash); err != nil {
				fmt.Println("insert file err:\n", err)
			}
		}
	})

	return flag
}
