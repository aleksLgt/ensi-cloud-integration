package recommendationQueryProducts

import (
	"context"

	"ensi-cloud-integration/internal/domain/recommendationQueryProductsDomain"
)

type (
	ensiCloudService interface {
		SearchRecommendationQueryProducts(ctx context.Context, request *recommendationQueryProductsDomain.SearchRecommendationQueryProductsRequest) (*recommendationQueryProductsDomain.SearchRecommendationQueryProductsResponse, error)
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

func (h *Handler) SearchRecommendationQueryProducts(ctx context.Context, request *recommendationQueryProductsDomain.SearchRecommendationQueryProductsRequest) (*recommendationQueryProductsDomain.SearchRecommendationQueryProductsResponse, error) {
	return h.ensiCloudService.SearchRecommendationQueryProducts(ctx, request)
}
