package categories

import "context"

type (
	ensiCloudService interface {
		IndexCategories(ctx context.Context) error
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

func (h *Handler) IndexCategories(ctx context.Context) error {
	return nil
}
