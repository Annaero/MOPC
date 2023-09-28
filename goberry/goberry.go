package main

import (
	"net/http"

	"github.com/Annaero/MOPC/goberry/router"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Get router with all routes.
	router.GetRoutes(r)

	// Run server instance.
	http.ListenAndServe(":3000", r)
}
