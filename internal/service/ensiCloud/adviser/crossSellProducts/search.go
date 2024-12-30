package crossSellProducts

import (
	"context"

	"ensi-cloud-integration/internal/domain"
)

type (
	ensiCloudService interface {
		SearchCrossSellProducts(
			ctx context.Context,
			request *domain.SearchCrossSellProductsRequest,
		) (*domain.SearchCrossSellProductsResponse, error)
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

func (h *Handler) SearchCrossSellProducts(
	ctx context.Context,
	request *domain.SearchCrossSellProductsRequest,
) (*domain.SearchCrossSellProductsResponse, error) {
	return h.ensiCloudService.SearchCrossSellProducts(ctx, request)
}
