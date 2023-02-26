package handler

import (
	"encoding/json"
	"net/http"

	"newsfeeder/platform/newsfeed"

	"gorm.io/gorm"
)

func NewsfeedGet(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var results []newsfeed.Item
		db.Find(&results) // Grab the whole data base as arrays of items
		json.NewEncoder(w).Encode(results)
	}
}
