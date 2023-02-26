package handler

import (
	"encoding/json"
	"net/http"
	"newsfeeder/platform/newsfeed"

	"gorm.io/gorm"
)

type newsfeedPostRequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		requestBody := newsfeedPostRequest{}
		json.NewDecoder(r.Body).Decode(&requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		db.Create(&item)
		json.NewEncoder(w).Encode("Added Post")
	}

}
