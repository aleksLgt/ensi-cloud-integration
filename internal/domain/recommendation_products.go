package domain

type (
	SearchRecommendationProductsRequest struct {
		Filter     filterByProductId   `json:"filter" validate:"nonnil"`
		Pagination paginationWithLimit `json:"pagination,omitempty"`
	}

	SearchRecommendationProductsResponse struct {
		Data struct {
			Products []string `json:"products"`
		} `json:"data"`
	}
)
