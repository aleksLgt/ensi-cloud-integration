package products

import (
	"context"

	"ensi-cloud-integration/internal/app/http/indexes/products"
)

type (
	ensiCloudService interface {
		IndexProducts(ctx context.Context, request *products.IndexProductsRequest) error
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

func (h *Handler) IndexProducts(ctx context.Context, request *products.IndexProductsRequest) error {
	return h.ensiCloudService.IndexProducts(ctx, request)
}
