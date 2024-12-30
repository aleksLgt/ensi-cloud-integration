package recommendationProducts

import (
	"context"

	"ensi-cloud-integration/internal/app/http/adviser/recommendationProducts"
	"ensi-cloud-integration/internal/clients/ensiCloud"
)

type (
	ensiCloudService interface {
		SearchRecommendationProducts(ctx context.Context, request *recommendationProducts.SearchRecommendationProductsRequest) (*ensiCloud.SearchRecommendationProductsResponse, error)
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

func (h *Handler) SearchRecommendationProducts(ctx context.Context, request *recommendationProducts.SearchRecommendationProductsRequest) (*ensiCloud.SearchRecommendationProductsResponse, error) {
	return h.ensiCloudService.SearchRecommendationProducts(ctx, request)
}
