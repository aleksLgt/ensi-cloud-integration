package recommendationProducts

import (
	"context"

	"ensi-cloud-integration/internal/domain"
)

type (
	ensiCloudService interface {
		SearchRecommendationProducts(
			ctx context.Context,
			request *domain.SearchRecommendationProductsRequest,
		) (*domain.SearchRecommendationProductsResponse, error)
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

func (h *Handler) SearchRecommendationProducts(
	ctx context.Context,
	request *domain.SearchRecommendationProductsRequest,
) (*domain.SearchRecommendationProductsResponse, error) {
	return h.ensiCloudService.SearchRecommendationProducts(ctx, request)
}
