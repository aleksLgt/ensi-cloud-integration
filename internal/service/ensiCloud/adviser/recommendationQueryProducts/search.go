package recommendationQueryProducts

import "context"

type (
	ensiCloudService interface {
		SearchRecommendedQueryProducts(ctx context.Context) error
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

func (h *Handler) SearchRecommendedQueryProducts(ctx context.Context) error {
	return nil
}
