package sqlite3

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	"litestream-sample-app/database"
)

type SQLite3 struct {
	conn *sql.DB
}

func (db *SQLite3) Exec(statement string, args ...interface{}) (database.Result, error) {
	result, err := db.conn.Exec(statement, args...)
	if err != nil {
		return &SQLite3Result{}, err
	}
	return &SQLite3Result{result: result}, nil
}

func (db *SQLite3) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := db.conn.Query(statement, args...)
	if err != nil {
		return &SQLite3Row{}, err
	}
	return &SQLite3Row{rows: rows}, nil
}

func (db *SQLite3) Close() error {
	db.conn.Close()
	return nil
}

type SQLite3Result struct {
	result sql.Result
}

func (r *SQLite3Result) LastInsertId() (int64, error) {
	return r.result.LastInsertId()
}

func (r *SQLite3Result) RowsAffected() (int64, error) {
	return r.result.RowsAffected()
}

type SQLite3Row struct {
	rows *sql.Rows
}

func (r *SQLite3Row) Scan(args ...interface{}) error {
	return r.rows.Scan(args...)
}

func (r *SQLite3Row) Next() bool {
	return r.rows.Next()
}

func (r *SQLite3Row) Close() error {
	return r.rows.Close()
}

func New(config *Config) (database.Database, error) {
	conn, err := sql.Open("sqlite3", config.DatabasePath)
	if err != nil {
		return &SQLite3{}, err
	}

	return &SQLite3{conn: conn}, nil
}
