package storage

import (
	"database/sql"
	"mime/multipart"

	_ "github.com/go-sql-driver/mysql"
)

const (
	TABLE  = "content"
	SCHEMA = `
	id CHAR(255) UNIQUE PRIMARY KEY NOT NULL,
	path CHAR(255) UNIQUE NOT NULL,
	mime CHAR(32) NOT NULL`
)

var (
	database *sql.DB
)

func Connect(address string) {
	var err error
	if database, err = sql.Open("mysql", address); err != nil {
		panic(err)
	}

	if err = database.Ping(); err != nil {
		panic(err)
	}

	if _, err = database.Exec("CREATE TABLE IF NOT EXISTS " + TABLE + "(" + SCHEMA + ")"); err != nil {
		panic(err)
	}

	database.SetMaxOpenConns(150)
}

func ReadPath(id string) (path string, exists bool, err error) {
	var statement string = "SELECT path FROM " + TABLE + " WHERE id=?"
	if err = database.QueryRow(statement, id).Scan(&path); err != nil {
		if err == sql.ErrNoRows {
			err = nil
		}

		return
	}

	exists = true
	return
}

func DeleteID(id string) (err error) {
	// TODO: should this be a const?
	var statement string = "DELETE FROM " + TABLE + " WHERE id=? LIMIT 1"
	_, err = database.Exec(statement, id)
	return
}

func CreateFile(id, mime string, file multipart.File) (err error) {
	var path string
	if path, err = WriteMultipartFile(id, file); err != nil {
		return
	}

	var statement string = "REPLACE INTO " + TABLE + " (id, mime, path) VALUES (?, ?, ?)"
	_, err = database.Exec(statement, id, mime, path)
	return
}
