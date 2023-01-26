package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
    "strconv"
)

var db *sql.DB

func SetDB() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   "root",
        Passwd: "wkfgkwk!@12",
        Net:    "tcp",
        Addr:   "192.168.0.9:3306",
        DBName: "data",
        AllowNativePasswords: true,
    }
    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}

func Query(sql string) string{
	rows, err := db.Query(sql)
	if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    result := "["
    cols, _ := rows.Columns()
    pointers := make([]interface{}, len(cols))
    container := make([]string, len(cols))
    for i, _ := range pointers {
        pointers[i] = &container[i]
    }

    for rows.Next() {
        

        if err := rows.Scan(pointers...); err != nil {
            fmt.Errorf("err : %v",  err)
        }
        result = result+"{"
        for inx, val := range container{
            result = result+"\""+cols[inx]+"\":\""+val+"\","
        }
        result = result[:len(result)-1]
        result = result+"},"

    }
    result = result[:len(result)-1]
    result = result+"]"
    return result

}

func Exec(sql string) string{
	result, err := db.Exec(sql)
	if err != nil {
        log.Fatal(err)
    }

	n, err := result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }

    return strconv.Itoa(int(n))
}


