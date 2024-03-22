package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LuanTruongPTIT/tutor-be/modules/server"
	"github.com/LuanTruongPTIT/tutor-be/pkg/dbs"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {

	dsn := os.Getenv("DBConnectionStr")
	fmt.Println("dsn", dsn)
	db, err := dbs.NewDatabase(dsn)
	httpSvr := server.NewServer(db)
	if err = httpSvr.Run(); err != nil {
		log.Fatal(err)
	}
}
