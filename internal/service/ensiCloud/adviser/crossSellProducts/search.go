package crossSellProducts

import (
	"context"

	"ensi-cloud-integration/internal/app/http/adviser/crossSellProducts"
	"ensi-cloud-integration/internal/clients/ensiCloud"
)

type (
	ensiCloudService interface {
		SearchCrossSellProducts(ctx context.Context, request *crossSellProducts.SearchCrossSellProductsRequest) (*ensiCloud.SearchCrossSellProductsResponse, error)
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

func (h *Handler) SearchCrossSellProducts(ctx context.Context, request *crossSellProducts.SearchCrossSellProductsRequest) (*ensiCloud.SearchCrossSellProductsResponse, error) {
	return h.ensiCloudService.SearchCrossSellProducts(ctx, request)
}
