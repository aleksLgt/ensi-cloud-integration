package catalog

import (
	"context"

	"ensi-cloud-integration/internal/app/http/catalog"
	"ensi-cloud-integration/internal/clients/ensiCloud"
)

type (
	ensiCloudService interface {
		SearchCatalog(ctx context.Context, request *catalog.SearchCatalogRequest) (*ensiCloud.SearchCatalogResponse, error)
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

func (h *Handler) SearchCatalog(ctx context.Context, request *catalog.SearchCatalogRequest) (*ensiCloud.SearchCatalogResponse, error) {
	return h.ensiCloudService.SearchCatalog(ctx, request)
}
