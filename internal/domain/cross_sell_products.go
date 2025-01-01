package domain

type (
	filterByProductId struct {
		ProductId string `json:"product_id" validate:"nonzero,nonnil"`
	}

	paginationWithLimit struct {
		Limit int `json:"limit,omitempty" validate:"max=50"`
	}

	SearchCrossSellProductsRequest struct {
		Filter     filterByProductId   `json:"filter" validate:"nonnil"`
		Pagination paginationWithLimit `json:"pagination,omitempty"`
	}

	SearchCrossSellProductsResponse struct {
		Data struct {
			Products []string `json:"products"`
		} `json:"data"`
		Errors []errorResponse `json:"errors,omitempty"`
		Meta   metaResponse    `json:"meta"`
	}
)
