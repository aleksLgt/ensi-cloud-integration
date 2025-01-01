package domain

type (
	indexCategoriesBody struct {
		Name      string   `json:"name" validate:"nonzero,nonnil"`
		URL       string   `json:"url,omitempty"`
		ParentIds []string `json:"parent_ids,omitempty"`
	}

	indexCategoriesAction struct {
		Action string              `json:"action"` // TODO custom rule
		Id     string              `json:"id" validate:"nonzero,nonnil"`
		Body   indexCategoriesBody `json:"body" validate:"nonnil"`
	}

	IndexCategoriesRequest struct {
		Actions []indexCategoriesAction `json:"actions" validate:"nonnil"`
	}

	IndexCategoriesResponse struct {
		Data   struct{}        `json:"data"`
		Errors []errorResponse `json:"errors,omitempty"`
		Meta   metaResponse    `json:"meta"`
	}
)
