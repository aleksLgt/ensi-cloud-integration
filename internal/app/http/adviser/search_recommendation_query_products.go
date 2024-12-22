package adviser

import (
	"context"
	"net/http"
)

type (
	searchRecommendationQueryProductsCommand interface {
		SearchRecommendationQueryProducts(ctx context.Context) error
	}

	SearchRecommendationQueryProductsHandler struct {
		name                                     string
		searchRecommendationQueryProductsCommand searchRecommendationQueryProductsCommand
	}

	searchRecommendationQueryProductsRequest struct {
	}
)

func NewSearchRecommendationQueryProductsHandler(command searchRecommendationQueryProductsCommand, name string) *SearchRecommendationQueryProductsHandler {
	return &SearchRecommendationQueryProductsHandler{
		name:                                     name,
		searchRecommendationQueryProductsCommand: command,
	}
}

func (h *SearchRecommendationQueryProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
