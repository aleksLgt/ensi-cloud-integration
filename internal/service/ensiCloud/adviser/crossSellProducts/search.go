package crossSellProducts

import "context"

type (
	ensiCloudService interface {
		SearchCrossSellProducts(ctx context.Context) error
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

func (h *Handler) SearchCrossSellProducts(ctx context.Context) error {
	return nil
}
