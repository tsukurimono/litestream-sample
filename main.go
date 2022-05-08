package main

import (
    "fmt"
    "time"

    "litestream-sample-app/database/sqlite3"
)

func main() {
    db, err := sqlite3.New()
    if err != nil {
		panic(err)
	}

	result, err := db.Exec("INSERT INTO userinfo(username, departname, created) values(?,?,?)", "astaxie", "研究開発部門", "2012-12-09")
    if err != nil {
		panic(err)
	}

    id, err := result.LastInsertId()
    if err != nil {
		panic(err)
	}

    fmt.Println(id)

    result, err = db.Exec("update userinfo set username=? where uid=?", "astaxieupdate", id)
    if err != nil {
		panic(err)
	}

    affect, err := result.RowsAffected()
    if err != nil {
		panic(err)
	}

    fmt.Println(affect)

    rows, err := db.Query("SELECT * FROM userinfo")
    if err != nil {
		panic(err)
	}

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created time.Time
        err = rows.Scan(&uid, &username, &department, &created)
        if err != nil {
		    panic(err)
	    }

        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }

    result, err = db.Exec("delete from userinfo where uid=?", id)
    if err != nil {
		panic(err)
	}

    affect, err = result.RowsAffected()
    if err != nil {
		panic(err)
	}

    fmt.Println(affect)

    //db.Close()
}
