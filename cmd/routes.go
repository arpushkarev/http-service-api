package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func route() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", answer)

	return mux
}
