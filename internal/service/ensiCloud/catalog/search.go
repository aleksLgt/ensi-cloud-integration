package catalog

import (
	"context"

	"ensi-cloud-integration/internal/domain"
)

type (
	ensiCloudService interface {
		SearchCatalog(ctx context.Context, request *domain.SearchCatalogRequest) (*domain.SearchCatalogResponse, error)
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

func (h *Handler) SearchCatalog(
	ctx context.Context,
	request *domain.SearchCatalogRequest,
) (*domain.SearchCatalogResponse, error) {
	return h.ensiCloudService.SearchCatalog(ctx, request)
}
