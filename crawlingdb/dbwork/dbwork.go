package main

import (
	
	db "dbwork/db"
	//web "dbwork/web"
	"log"

)


func main() {
	db.SetDB()

	re := db.Perser()
	log.Println(re)

	//web.Router()

}
