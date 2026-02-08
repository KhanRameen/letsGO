package handler

import (
	"fmt"

	"github.com/go-chi/chi/middleware"
	chimiddle "github.com/go-chi/chi/middleware"
	"builds/api-with-go/internal/middleware"
)

func Handler( r *chi.Mux) { 
	r.Use(chimiddle.StripSlashes) //use the middleware function for all routes
	//strip slashes from the URL path, so that /api and /api/ are treated the same

	r.Route("/account", func(router chi.Router) { //creates a router group for /account
		
		router.Use(middleware.Authorization) //middlware for /account

		router.Get("/coin", GetCoinBalance) //route for /account/coin


	})
}
