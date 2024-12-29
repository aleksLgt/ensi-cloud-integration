package categories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	http2 "ensi-cloud-integration/internal/app/http"

	"gopkg.in/validator.v2"
)

type (
	indexCategoriesCommand interface {
		IndexCategories(ctx context.Context, request *IndexCategoriesRequest) error
	}

	IndexCategoriesHandler struct {
		name                   string
		indexCategoriesCommand indexCategoriesCommand
	}

	bodyRequest struct {
		Name      string   `json:"name" validate:"nonzero,nonnil"`
		URL       string   `json:"url"`
		ParentIds []string `json:"parent_ids"`
	}

	actionRequest struct {
		Action string      `json:"action"` // TODO custom rule
		Id     string      `json:"id" validate:"nonzero,nonnil"`
		Body   bodyRequest `json:"body" validate:"nonnil"`
	}

	IndexCategoriesRequest struct {
		Actions []actionRequest `json:"actions" validate:"nonnil"`
	}
)

func NewIndexCategoriesHandler(command indexCategoriesCommand, name string) *IndexCategoriesHandler {
	return &IndexCategoriesHandler{
		name:                   name,
		indexCategoriesCommand: command,
	}
}

func (h *IndexCategoriesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	var (
		request *IndexCategoriesRequest
		err     error
	)

	if request, err = h.getRequestData(r); err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err = validator.Validate(request); err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	err = h.indexCategoriesCommand.IndexCategories(ctx, request)
	if err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	http2.GetSuccessResponse(w)
}

func (_ *IndexCategoriesHandler) getRequestData(r *http.Request) (*IndexCategoriesRequest, error) {
	request := &IndexCategoriesRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
