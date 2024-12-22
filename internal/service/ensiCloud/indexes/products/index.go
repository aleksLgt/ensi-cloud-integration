package products

import "context"

type (
	ensiCloudService interface {
		IndexProducts(ctx context.Context) error
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

func (h *Handler) IndexProducts(ctx context.Context) error {
	return nil
}
