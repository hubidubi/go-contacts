package main

import (
	"flag"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"main/app"
	"net/http"
	"time"
)

func main() {
	port := flag.String("port", "8080", "port to use")
	timeout := flag.String("timeout", "10", "timeout value")
	duration, _ := time.ParseDuration(*timeout + "s")
	dsn := flag.String("dsn", "gocontacts:gocontacts@tcp(localhost:3306)/gocontacts?charset=utf8&parseTime=true", "DSN to use")
	flag.Parse()
	log.Printf("Timeout set to %ss", *timeout)

	if DB, err := gorm.Open(mysql.Open(*dsn), &gorm.Config{}); err == nil {
		log.Print("DB connected: " + *dsn)
		DB.AutoMigrate(&app.Contact{})
		app.DB = DB
	} else {
		log.Fatal("failed to connect database " + *dsn)
		return
	}
	log.Print("Go contacts is running on port " + *port)
	srv := &http.Server{
		Handler:      app.GetRouter(),
		Addr:         ":" + *port,
		ReadTimeout:  duration,
		WriteTimeout: duration,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
