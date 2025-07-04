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

func QueryFileByPath(filepath string) FileMark {
	var file FileMark
	dbfile := GetMarkTable(true)

	Connection(dbfile, func(db *sql.DB) {
		rows, err := db.Query(QUERY_FILE_BY_PATH, filepath)
		if err != nil {
			fmt.Println("query files err:\n", err)
		}
		defer rows.Close()
		if rows.Next() {
			err := rows.Scan(&file.ID, &file.FilePath, &file.Dir, &file.Sha256, &file.LastModify, &file.ModifyAt, &file.CreatedAt)
			if err != nil {
				fmt.Println("scan files err:\n", err)
			}
		}
	})

	return file
}

func QueryFileBySHA256(sha256 string) FileMark {
	var file FileMark
	dbfile := GetMarkTable(true)

	Connection(dbfile, func(db *sql.DB) {
		rows, err := db.Query(QUERY_FILE_BY_SHA256, sha256)
		if err != nil {
			fmt.Println("query files err:\n", err)
		}
		defer rows.Close()
		if rows.Next() {
			err := rows.Scan(&file.ID, &file.FilePath, &file.Dir, &file.Sha256, &file.CreatedAt, &file.ModifyAt)
			if err != nil {
				fmt.Println("scan files err:\n", err)
			}
		}
	})

	return file
}

func QueryFilesByDir(dir string) []FileMark {
	result := make([]FileMark, 0)
	dbfile := GetMarkTable(true)

	Connection(dbfile, func(db *sql.DB) {
		rows, err := db.Query(QUERY_ALL_FILES_BY_DIR, dir)
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
		info, _ := os.Stat(path)
		last_modify := info.ModTime().Local()
		existing := QueryFileByPath(path)
		hash := "/"

		if existing.ID != 0 {
			if !existing.ModifyAt.Equal(last_modify) {
				hash = existing.Sha256

				if !info.IsDir() {
					if sha, err := CalculateFileSHA256(path); err != nil {
						fmt.Println("calculate file sha256 failure\n", err)
					} else {
						hash = sha
					}
				}

				if _, err := db.Exec(CHANGE_FILE_SQL, path, filepath.Dir(path), hash, last_modify, existing.ID); err != nil {
					fmt.Println("change file info err:\n", err)
				} else {
					flag = true
				}
			}
			return
		}

		if _, err := db.Exec(INSERT_FILE_SQL, path, filepath.Dir(path), hash, last_modify); err != nil {
			fmt.Println("insert file err:\n", err)
		}
	})

	return flag
}

func ChangeFileInfo(fileInfo *FileMark) bool {
	dbfile := GetMarkTable(true)
	flag := false

	Connection(dbfile, func(db *sql.DB) {
		if _, err := db.Exec(CHANGE_FILE_SQL, fileInfo.FilePath, fileInfo.Dir, fileInfo.Sha256, fileInfo.ID); err != nil {
			fmt.Println("change file info err:\n", err)
		} else {
			flag = true
		}
	})

	return flag
}
