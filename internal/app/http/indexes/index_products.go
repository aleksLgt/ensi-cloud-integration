package indexes

import (
	"context"
	"net/http"
)

type (
	indexProductsCommand interface {
		IndexProducts(ctx context.Context) error
	}

	IndexProductsHandler struct {
		name                 string
		indexProductsCommand indexProductsCommand
	}

	indexProductsRequest struct {
	}
)

func NewIndexProductsHandler(command indexProductsCommand, name string) *IndexProductsHandler {
	return &IndexProductsHandler{
		name:                 name,
		indexProductsCommand: command,
	}
}

func (h *IndexProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
