package categories

import (
	"context"

	"ensi-cloud-integration/internal/domain"
)

type (
	ensiCloudService interface {
		IndexCategories(ctx context.Context, request *domain.IndexCategoriesRequest) (*domain.IndexCategoriesResponse, error)
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

func (h *Handler) IndexCategories(ctx context.Context, request *domain.IndexCategoriesRequest) (*domain.IndexCategoriesResponse, error) {
	return h.ensiCloudService.IndexCategories(ctx, request)
}
