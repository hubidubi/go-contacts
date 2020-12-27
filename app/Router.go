package app

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func GetRouter() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Go Contacts"))
	})
	muxRouter.Handle("/metrics", promhttp.Handler()).Methods("GET")
	contactsRouter := muxRouter.PathPrefix("/contacts").Subrouter()
	contactsRouter.HandleFunc("", ListContacts).Methods("GET")
	contactsRouter.HandleFunc("/{id}", GetContact).Methods("GET")
	contactsRouter.HandleFunc("", AddContact).Methods("POST")
	contactsRouter.HandleFunc("/{id}", UpdateContact).Methods("PUT")
	contactsRouter.HandleFunc("/{id}", DeleteContact).Methods("DELETE")
	contactsRouter.HandleFunc("/name/{name}", SearchContact).Methods("GET")

	return muxRouter
}
