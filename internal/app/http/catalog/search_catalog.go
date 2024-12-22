package catalog

import (
	"context"
	"net/http"
)

type (
	searchCatalogCommand interface {
		SearchCatalog(ctx context.Context) error
	}

	SearchCatalogHandler struct {
		name                 string
		searchCatalogCommand searchCatalogCommand
	}

	searchCatalogRequest struct {
	}
)

func NewSearchCatalogHandler(command searchCatalogCommand, name string) *SearchCatalogHandler {
	return &SearchCatalogHandler{
		name:                 name,
		searchCatalogCommand: command,
	}
}

func (h *SearchCatalogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
