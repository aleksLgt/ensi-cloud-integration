package adviser

import (
	"context"
	"net/http"
)

type (
	searchRecommendationProductsCommand interface {
		SearchRecommendationProducts(ctx context.Context) error
	}

	SearchRecommendationProductsHandler struct {
		name                                string
		searchRecommendationProductsCommand searchRecommendationProductsCommand
	}

	searchRecommendationProductsRequest struct {
	}
)

func NewSearchRecommendationProductsHandler(command searchRecommendationProductsCommand, name string) *SearchRecommendationProductsHandler {
	return &SearchRecommendationProductsHandler{
		name:                                name,
		searchRecommendationProductsCommand: command,
	}
}

func (h *SearchRecommendationProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}
