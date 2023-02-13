package search

import (
	"github.com/go-chi/chi"
)

// ZincSearchRoute function that return
func ZincSearchRoute() *chi.Mux {
	zincSearchRoute := chi.NewRouter()
	zincSearchRoute.Post("/{index}/get_all", GetAll)
	zincSearchRoute.Post("/{index}/search_emails", GetSearchEmails)

	return zincSearchRoute
}
