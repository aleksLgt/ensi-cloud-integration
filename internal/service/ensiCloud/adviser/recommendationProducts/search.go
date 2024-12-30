package recommendationProducts

import (
	"context"

	"ensi-cloud-integration/internal/domain/recommendationProductsDomain"
)

type (
	ensiCloudService interface {
		SearchRecommendationProducts(ctx context.Context, request *recommendationProductsDomain.SearchRecommendationProductsRequest) (*recommendationProductsDomain.SearchRecommendationProductsResponse, error)
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

func (h *Handler) SearchRecommendationProducts(ctx context.Context, request *recommendationProductsDomain.SearchRecommendationProductsRequest) (*recommendationProductsDomain.SearchRecommendationProductsResponse, error) {
	return h.ensiCloudService.SearchRecommendationProducts(ctx, request)
}
