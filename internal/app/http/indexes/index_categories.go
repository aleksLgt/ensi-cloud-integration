package indexes

import (
	"context"
	"net/http"
)

type (
	indexCategoriesCommand interface {
		IndexCategories(ctx context.Context) error
	}

	IndexCategoriesHandler struct {
		name                   string
		indexCategoriesCommand indexCategoriesCommand
	}

	indexCategoriesRequest struct {
	}
)

func NewIndexCategoriesHandler(command indexCategoriesCommand, name string) *IndexCategoriesHandler {
	return &IndexCategoriesHandler{
		name:                   name,
		indexCategoriesCommand: command,
	}
}

func (h *IndexCategoriesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
