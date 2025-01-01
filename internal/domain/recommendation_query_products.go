package domain

type (
	filter struct {
		Query string `json:"query" validate:"nonzero,nonnil"`
	}

	SearchRecommendationQueryProductsRequest struct {
		Filter     filter              `json:"filter" validate:"nonnil"`
		Pagination paginationWithLimit `json:"pagination,omitempty"`
	}

	SearchRecommendationQueryProductsResponse struct {
		Data struct {
			Products []string `json:"products"`
		} `json:"data"`
		Errors []errorResponse `json:"errors,omitempty"`
		Meta   metaResponse    `json:"meta"`
	}
)
