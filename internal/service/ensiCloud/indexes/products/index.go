package products

import (
	"context"

	"ensi-cloud-integration/internal/domain"
)

type (
	ensiCloudService interface {
		IndexProducts(ctx context.Context, request *domain.IndexProductsRequest) (*domain.IndexProductsResponse, error)
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

func (h *Handler) IndexProducts(ctx context.Context, request *domain.IndexProductsRequest) (*domain.IndexProductsResponse, error) {
	return h.ensiCloudService.IndexProducts(ctx, request)
}
