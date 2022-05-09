package database

import (
	"bytes"
	"io/ioutil"
	"strings"
)

type Database interface {
	Exec(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
	Close() error
}

type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

func Initialize(db Database, initSQLPath string) error {
	content, err := ioutil.ReadFile(initSQLPath)
	if err != nil {
		return err
	}
	sqlText := bytes.NewBuffer(content).String()
	sqlSlice := strings.Split(sqlText, ";\n")
	for _, sql := range sqlSlice {
		if strings.TrimSpace(sql) == "" {
			continue
		}
		_, err := db.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}
