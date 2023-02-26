package main

import (
	"log"
	"net/http"

	"newsfeeder/platform/newsfeed"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("../platform/newsfeed/newsfeed.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&newsfeed.Item{})

	host := "127.0.0.1:5000"
	if err := http.ListenAndServe(host, httpHandler(db)); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}
}
