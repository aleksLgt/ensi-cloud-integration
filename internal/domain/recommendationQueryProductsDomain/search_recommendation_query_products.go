package recommendationQueryProductsDomain

type (
	filter struct {
		Query string `json:"query" validate:"nonzero,nonnil"`
	}

	pagination struct {
		Limit int `json:"limit,omitempty" validate:"max=50"`
	}

	SearchRecommendationQueryProductsRequest struct {
		Filter     filter     `json:"filter" validate:"nonnil"`
		Pagination pagination `json:"pagination,omitempty"`
	}

	SearchRecommendationQueryProductsResponse struct {
		Data struct {
			Products []string `json:"products"`
		} `json:"data"`
	}
)
