package catalog

import "context"

type (
	ensiCloudService interface {
		SearchCatalog(ctx context.Context) error
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

func (h *Handler) SearchCatalog(ctx context.Context) error {
	return nil
}
