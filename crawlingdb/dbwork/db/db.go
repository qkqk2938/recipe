package db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
    "strconv"
    "strings"
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


func Query(sql string) []map[string]string{
    log.Println(sql)
	rows, err := db.Query(sql)
	if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    cols, _ := rows.Columns()
    pointers := make([]interface{}, len(cols))
    container := make([]string, len(cols))
    result := make([]map[string]string,0)
    for i, _ := range pointers {
        pointers[i] = &container[i]
    }

    for rows.Next() {
        if err := rows.Scan(pointers...); err != nil {
            fmt.Errorf("err : %v",  err)
        }
        item := make(map[string]string)
        for i, v := range cols {
            item[v] = container[i]
        }
        result = append(result, item)
    }
    return result

}


func JsonQuery(sql string) string{
    log.Println(sql)
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
            result = result+"\""+cols[inx]+"\":\""+strings.Replace(val, "\"", "\\\"",-1)+"\","
        }
        result = result[:len(result)-1]
        result = result+"},"

    }
    result = result[:len(result)-1]
    result = result+"]"
    return result

}

func Exec(sql string) string{
    log.Println(sql)
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

func ExecGetLastID(sql string) string{
    log.Println(sql)
	result, err := db.Exec(sql)
	if err != nil {
        log.Fatal(err)
    }

	_, err = result.RowsAffected()
    if err != nil {
        log.Fatal(err)
    }
    lastID, _ := result.LastInsertId()
    return strconv.Itoa(int(lastID))
}