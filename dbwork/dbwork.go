package main

import (
	
	db "dbwork/db"
	web "dbwork/web"


)


func main() {
	db.SetDB()

	//db.Perser()


	web.Router()

}
