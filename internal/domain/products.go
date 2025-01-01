package domain

type (
	location struct {
		Id    string `json:"id" validate:"nonzero,nonnil"`
		Price int    `json:"price,omitempty" validate:"nonzero,nonnil"`
	}

	indexProductsBody struct {
		Name        string     `json:"name" validate:"nonzero,nonnil"`
		URL         string     `json:"url,omitempty"`
		CategoryIds []string   `json:"category_ids" validate:"nonnil,min=1"`
		Brand       string     `json:"brand,omitempty"`
		VendorCode  string     `json:"vendor_code" validate:"nonzero,nonnil"`
		Barcodes    []string   `json:"barcodes,omitempty"`
		Description string     `json:"description,omitempty"`
		Picture     string     `json:"picture,omitempty"`
		Country     string     `json:"country,omitempty"`
		GroupIds    []string   `json:"group_ids,omitempty"`
		Locations   []location `json:"locations" validate:"nonnil"`
		Properties  []property `json:"properties,omitempty" validate:"nonnil"`
	}

	indexProductsAction struct {
		Action string            `json:"action" validate:"actionTypeRule"`
		Id     string            `json:"id" validate:"nonzero,nonnil"`
		Body   indexProductsBody `json:"body" validate:"nonnil"`
	}

	IndexProductsRequest struct {
		Actions []indexProductsAction `json:"actions" validate:"nonnil"`
	}

	IndexProductsResponse struct {
		Data   struct{}        `json:"data"`
		Errors []errorResponse `json:"errors,omitempty"`
		Meta   metaResponse    `json:"meta"`
	}
)
