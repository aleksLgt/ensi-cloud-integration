package recommendationQueryProducts

import (
	"context"

	"ensi-cloud-integration/internal/app/http/adviser/recommendationQueryProducts"
)

type (
	ensiCloudService interface {
		SearchRecommendationQueryProducts(ctx context.Context, request *recommendationQueryProducts.SearchRecommendationQueryProductsRequest) ([]byte, error)
	}

	Handler struct {
		ensiCloudService ensiCloudService
	}
)

func New(ensiCloudService ensiCloudService) *Handler {
	return &Handler{
		ensiCloudService: ensiCloudService,
	}
}

func (h *Handler) SearchRecommendationQueryProducts(ctx context.Context, request *recommendationQueryProducts.SearchRecommendationQueryProductsRequest) ([]byte, error) {
	return h.ensiCloudService.SearchRecommendationQueryProducts(ctx, request)
}
