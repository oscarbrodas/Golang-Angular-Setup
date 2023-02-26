package main

import (
	"net/http"
	"os"

	"newsfeeder/httpd/handler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func httpHandler(db *gorm.DB) http.Handler {
	router := mux.NewRouter()
	apiRoute := router.PathPrefix("/api").Subrouter()

	apiRoute.HandleFunc("/ping", handler.PingGet()).Methods("GET")
	apiRoute.HandleFunc("/newsfeed", handler.NewsfeedGet(db)).Methods("GET")
	apiRoute.HandleFunc("/newsfeed", handler.NewsfeedPost(db)).Methods("POST")

	router.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	return handlers.LoggingHandler(os.Stdout,
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
				"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
				"Cache-Control", "Content-Range", "Range"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"http://localhost:8080"}),
			handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
				"X-Requested-With", "If-Modified-Since", "Cache-Control",
				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
			handlers.MaxAge(86400),
		)(router))
}
