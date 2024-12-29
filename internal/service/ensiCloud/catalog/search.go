package catalog

import (
	"context"

	"ensi-cloud-integration/internal/app/http/catalog"
)

type (
	ensiCloudService interface {
		SearchCatalog(ctx context.Context, request *catalog.SearchCatalogRequest) ([]byte, error)
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

func (h *Handler) SearchCatalog(ctx context.Context, request *catalog.SearchCatalogRequest) ([]byte, error) {
	return h.ensiCloudService.SearchCatalog(ctx, request)
}
