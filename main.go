package main

import (
	"fmt"

	"litestream-sample-app/database"
	"litestream-sample-app/database/sqlite3"
)

func main() {
	db, err := sqlite3.New(&sqlite3.Config{DatabasePath: "/var/datafiles/db"})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = database.Initialize(db, "/init/init.sql")
	if err != nil {
		panic(err)
	}

	result, err := db.Exec("INSERT INTO fruits (name, color) VALUES (?, ?);", "apple", "red")
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)

	result, err = db.Exec("update fruits set color=? where name=?", "blue", "apple")
	if err != nil {
		panic(err)
	}

	affect, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(affect)

	rows, err := db.Query("SELECT * FROM fruits")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var name string
		var color string
		err = rows.Scan(&name, &color)
		if err != nil {
			panic(err)
		}

		fmt.Println(name)
		fmt.Println(color)
	}

	result, err = db.Exec("delete from fruits where name=?", "apple")
	if err != nil {
		panic(err)
	}

	affect, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(affect)
}
