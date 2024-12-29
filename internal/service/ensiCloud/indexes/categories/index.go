package categories

import (
	"context"

	"ensi-cloud-integration/internal/app/http/indexes/categories"
)

type (
	ensiCloudService interface {
		IndexCategories(ctx context.Context, request *categories.IndexCategoriesRequest) error
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

func (h *Handler) IndexCategories(ctx context.Context, request *categories.IndexCategoriesRequest) error {
	return h.ensiCloudService.IndexCategories(ctx, request)
}
