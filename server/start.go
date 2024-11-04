package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"pixel-haven/internal/config"
)

func Start(cxt context.Context, conf *config.Config) {
	router := chi.NewRouter()
	router.Group(func(router chi.Router) {
		router.Use(Api(conf))
		router.Use(Logger)
		RegisterRouters(router)
	})
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Println(err)
	}
}
