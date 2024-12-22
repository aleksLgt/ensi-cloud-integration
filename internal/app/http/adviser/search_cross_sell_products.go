package adviser

import (
	"context"
	"net/http"
)

type (
	searchCrossSellProductsCommand interface {
		SearchCrossSellProducts(ctx context.Context) error
	}

	SearchCrossSellProductsHandler struct {
		name                           string
		searchCrossSellProductsCommand searchCrossSellProductsCommand
	}

	searchCrossSellProductsRequest struct {
	}
)

func NewSearchCrossSellProductsHandler(command searchCrossSellProductsCommand, name string) *SearchCrossSellProductsHandler {
	return &SearchCrossSellProductsHandler{
		name:                           name,
		searchCrossSellProductsCommand: command,
	}
}

func (h *SearchCrossSellProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
