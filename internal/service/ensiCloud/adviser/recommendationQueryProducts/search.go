package recommendationQueryProducts

import (
	"context"

	"ensi-cloud-integration/internal/app/http/adviser/recommendationQueryProducts"
	"ensi-cloud-integration/internal/clients/ensiCloud"
)

type (
	ensiCloudService interface {
		SearchRecommendationQueryProducts(ctx context.Context, request *recommendationQueryProducts.SearchRecommendationQueryProductsRequest) (*ensiCloud.SearchRecommendationQueryProductsResponse, error)
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

func (h *Handler) SearchRecommendationQueryProducts(ctx context.Context, request *recommendationQueryProducts.SearchRecommendationQueryProductsRequest) (*ensiCloud.SearchRecommendationQueryProductsResponse, error) {
	return h.ensiCloudService.SearchRecommendationQueryProducts(ctx, request)
}
