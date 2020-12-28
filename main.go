package main

import (
	"flag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main/app"
	"net/http"
)

func main() {
	port := flag.String("port", "8080", "port to use")
	dsn := flag.String("dsn", "gocontacts:gocontacts@tcp(localhost:3306)/gocontacts?charset=utf8&parseTime=true", "DSN to use")
	flag.Parse()

	if DB, err := gorm.Open(mysql.Open(*dsn), &gorm.Config{}); err == nil {
		DB.AutoMigrate(&app.Contact{})
		app.DB = DB
	} else {
		log.Fatal("failed to connect database " + *dsn)
		return
	}
	log.Print("Go contacts is running on port " + *port)
	err := http.ListenAndServe(":"+*port, app.GetRouter())
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
