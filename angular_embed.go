package main

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed client
var assests embed.FS

func getFrontend() http.FileSystem {
	if subAsset, err := fs.Sub(assests, "dist"); err == nil {
		return http.FS(subAsset)
	}

	panic("Failed to load assests")
}

var _angularHandler = http.FileServer(getFrontend())

var AngularHandlerEmbed = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=604800")
		_angularHandler.ServeHTTP(w, r)
	},
)
