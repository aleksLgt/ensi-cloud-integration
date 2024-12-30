package catalog

import (
	"context"

	"ensi-cloud-integration/internal/domain/catalogDomain"
)

type (
	ensiCloudService interface {
		SearchCatalog(ctx context.Context, request *catalogDomain.SearchCatalogRequest) (*catalogDomain.SearchCatalogResponse, error)
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

func (h *Handler) SearchCatalog(ctx context.Context, request *catalogDomain.SearchCatalogRequest) (*catalogDomain.SearchCatalogResponse, error) {
	return h.ensiCloudService.SearchCatalog(ctx, request)
}
