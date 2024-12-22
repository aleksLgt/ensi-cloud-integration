package recommendationProducts

import "context"

type (
	ensiCloudService interface {
		SearchRecommendedProducts(ctx context.Context) error
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

func (h *Handler) SearchRecommendedProducts(ctx context.Context) error {
	return nil
}
