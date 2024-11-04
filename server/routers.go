package server

import (
	"github.com/go-chi/chi/v5"
	"pixel-haven/internal/config"
)

func RegisterRouters(r chi.Router) {
	r.Route(config.ApiUri+"/photos", func(r chi.Router) {
		//r.Post("/", api.UploadPhoto)
	})
}
