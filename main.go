package main

import (
	"log"
	"os"
	"strconv"
	"url-shortener/app"
)

func main() {
	db_addr := os.Getenv("DATABASE_URL")
	db_pass := os.Getenv("DATABASE_PASS")
	db_id, err := strconv.Atoi(os.Getenv("DATABASE_ID"))
	if err != nil {
		log.Fatal(err, " -- Cannot cast DATABASE_ID to int")
	}
	port := os.Getenv("PORT")

	app := app.NewServer(db_addr, db_pass, db_id)
	app.Run(port)
}
