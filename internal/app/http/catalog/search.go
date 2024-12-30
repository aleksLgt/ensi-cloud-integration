package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/validator.v2"

	http2 "ensi-cloud-integration/internal/app/http"
)

type (
	searchCatalogCommand interface {
		SearchCatalog(ctx context.Context, request *SearchCatalogRequest) ([]byte, error)
	}

	SearchCatalogHandler struct {
		name                 string
		searchCatalogCommand searchCatalogCommand
	}

	property struct {
		Name   string   `json:"name" validate:"nonzero,nonnil"`
		Values []string `json:"values" validate:"nonnil,min=1"`
	}

	filter struct {
		LocationId  string     `json:"location_id" validate:"nonzero,nonnil"`
		Query       string     `json:"query" validate:"nonzero,nonnil"`
		CategoryIds []string   `json:"category_ids,omitempty"`
		Brands      []string   `json:"brands,omitempty"`
		Countries   []string   `json:"countries,omitempty"`
		Properties  []property `json:"properties,omitempty"`
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
)

func NewSearchCatalogHandler(command searchCatalogCommand, name string) *SearchCatalogHandler {
	return &SearchCatalogHandler{
		name:                 name,
		searchCatalogCommand: command,
	}
}

func (h *SearchCatalogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	var (
		request *SearchCatalogRequest
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

	response, err := h.searchCatalogCommand.SearchCatalog(ctx, request)
	if err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	http2.GetSuccessResponseWithBody(w, response)
}

func (_ *SearchCatalogHandler) getRequestData(r *http.Request) (*SearchCatalogRequest, error) {
	request := &SearchCatalogRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
