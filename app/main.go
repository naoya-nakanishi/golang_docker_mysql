package main

import (
	"database/sql"
	"log"
)

type User struct {
	ID int
	Name string
	Password string
}
const (
	// DriverName ドライバ名(mysql固定)
	DriverName = "mysql"
	// DataSourceName user:password@tcp(container-name:port)/dbname
	DataSourceName = "root:golang@tcp(mysql-container:3306)/golang_db"
)


var usr = make(map[int]User)

func main() {
	// database
  // データベースへの接続 ①
	db, dbErr := sql.Open(DriverName, DataSourceName)
	if dbErr != nil {
		log.Print("error connecting to database:", dbErr)
	}
    defer db.Close()
    // usersテーブルの全てのレコードを取得するクエリの実行 ②
    rows, queryErr := db.Query("SELECT * FROM users")
    if queryErr != nil {
        log.Print("query error :", queryErr)
    }
    defer rows.Close()
    // ループを回してrowsからScanでデータを取得する。 ③
    for rows.Next() {
        var u User
        if err := rows.Scan(&u.ID, &u.Name, &u.Password); err != nil {
            log.Print(err)
        }
        usr[u.ID] = User{
            ID:       u.ID,
            Name:     u.Name,
            Password: u.Password,
        }
    }
}