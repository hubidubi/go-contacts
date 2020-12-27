package main

import (
	"log"
	"main/app"
	"main/app/config"
	"net/http"
)

func main() {
	if err := config.LoadConfig("./config"); err != nil {
		log.Fatalf("invalid application configuration: %s", err)
	}

	//loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	//http.ListenAndServe(":8000", loggedRouter)
	err := http.ListenAndServe(config.Config.Host+":"+config.Config.Port, app.GetRouter())
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
