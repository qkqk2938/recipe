package main

import (
	"log"
	db "dbwork/db"

)


func main() {
	db.SetDB()
	data := db.SelectBase(nil)
	log.Println(data)

}
