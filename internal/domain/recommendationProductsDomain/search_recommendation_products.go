package recommendationProductsDomain

type (
	filter struct {
		ProductId string `json:"product_id" validate:"nonzero,nonnil"`
	}

	pagination struct {
		Limit int `json:"limit,omitempty" validate:"max=50"`
	}

	SearchRecommendationProductsRequest struct {
		Filter     filter     `json:"filter" validate:"nonnil"`
		Pagination pagination `json:"pagination,omitempty"`
	}

	SearchRecommendationProductsResponse struct {
		Data struct {
			Products []string `json:"products"`
		} `json:"data"`
	}
)
