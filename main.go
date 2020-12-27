package main

import (
	"log"
	"main/app"
	"net/http"
)

func main() {
	if err := app.LoadConfig("./config"); err != nil {
		log.Fatalf("invalid application configuration: %s", err)
	}

	err := http.ListenAndServe(app.Config.Host, app.GetRouter())
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
