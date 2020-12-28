package app

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func GetRouter() *mux.Router {
	muxRouter := mux.NewRouter()
	muxRouter.Handle("/", handlers.LoggingHandler(log.Writer(), http.HandlerFunc(index))).Methods("GET")
	muxRouter.Handle("/metrics", handlers.LoggingHandler(log.Writer(), promhttp.Handler())).Methods("GET")
	contactsRouter := muxRouter.PathPrefix("/contacts").Subrouter()
	contactsRouter.Handle("/", handlers.LoggingHandler(log.Writer(), http.HandlerFunc(ListContacts))).Methods("GET")
	contactsRouter.Handle("/{id}", handlers.LoggingHandler(log.Writer(), http.HandlerFunc(GetContact))).Methods("GET")
	contactsRouter.Handle("", handlers.LoggingHandler(log.Writer(), http.HandlerFunc(AddContact))).Methods("POST")
	contactsRouter.Handle("/{id}", handlers.LoggingHandler(log.Writer(), http.HandlerFunc(UpdateContact))).Methods("PUT")
	contactsRouter.Handle("/{id}", handlers.LoggingHandler(log.Writer(), http.HandlerFunc(DeleteContact))).Methods("DELETE")
	contactsRouter.Handle("/name/{name}", handlers.LoggingHandler(log.Writer(), http.HandlerFunc(SearchContact))).Methods("GET")

	return muxRouter
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go Contacts"))
}
