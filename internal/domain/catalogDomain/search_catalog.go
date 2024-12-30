package catalogDomain

type (
	Property struct {
		Name   string   `json:"name" validate:"nonzero,nonnil"`
		Values []string `json:"values" validate:"nonnil,min=1"`
	}

	filter struct {
		LocationId  string     `json:"location_id" validate:"nonzero,nonnil"`
		Query       string     `json:"query" validate:"nonzero,nonnil"`
		CategoryIds []string   `json:"category_ids,omitempty"`
		Brands      []string   `json:"brands,omitempty"`
		Countries   []string   `json:"countries,omitempty"`
		Properties  []Property `json:"properties,omitempty"`
	}

	pagination struct {
		LimitProducts   int `json:"limit_products,omitempty" validate:"max=1000"`
		OffsetProducts  int `json:"offset_products,omitempty"`
		LimitCategories int `json:"limit_categories,omitempty" validate:"max=100"`
	}

	SearchCatalogRequest struct {
		IsFastResult bool       `json:"is_fast_result,omitempty"`
		Include      []string   `json:"include"`        // TODO custom rule
		Sort         string     `json:"sort,omitempty"` // TODO custom rule
		Filter       filter     `json:"filter" validate:"nonnil"`
		Pagination   pagination `json:"pagination" validate:"nonnil"`
	}

	SearchCatalogResponse struct {
		Data struct {
			TotalProducts   int      `json:"total_products"`
			Products        []string `json:"products"`
			TotalCategories int      `json:"total_categories"`
			Categories      []string `json:"categories"`
			Correction      string   `json:"correction"`
			ProductHints    []struct {
				Word string `json:"word"`
				Hint string `json:"hint"`
			} `json:"product_hints"`
			Filters []struct {
				Name   string `json:"name"`
				Code   string `json:"code"`
				Values []struct {
					Id   string `json:"id"`
					Name string `json:"name"`
				} `json:"values"`
			} `json:"filters"`
		} `json:"data"`
	}
)
