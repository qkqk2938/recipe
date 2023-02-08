package main

import (
	
	crawling "crawling_parser/crawling"
	db "crawling_parser/db"
	"log"
	"os"

)


func main() {

	runType := "all"
	if len(os.Args) >= 2{
		runType = os.Args[1]
	}
	log.Println("main : start")

	db.SetDB()
	log.Println("main : DB SET!")

	if runType == "crawling" || runType == "all" {
		crawling.GoCrawling()
		log.Println("main : crawling!")

	}
	if runType == "parser" || runType == "all" {
		re := db.Parser()
		log.Println("main : insert DB!")

		log.Println(re)
	}
	//web.Router()

}
