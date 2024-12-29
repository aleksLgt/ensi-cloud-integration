package crossSellProducts

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/validator.v2"

	http2 "ensi-cloud-integration/internal/app/http"
)

type (
	searchCrossSellProductsCommand interface {
		SearchCrossSellProducts(ctx context.Context, request *SearchCrossSellProductsRequest) ([]byte, error)
	}

	SearchCrossSellProductsHandler struct {
		name                           string
		searchCrossSellProductsCommand searchCrossSellProductsCommand
	}

	filterRequest struct {
		ProductId string `json:"product_id" validate:"nonzero,nonnil"`
	}

	paginationRequest struct {
		Limit int `json:"limit" validate:"max=50"`
	}

	SearchCrossSellProductsRequest struct {
		Filter     filterRequest     `json:"filter" validate:"nonnil"`
		Pagination paginationRequest `json:"pagination"`
	}
)

func NewSearchCrossSellProductsHandler(command searchCrossSellProductsCommand, name string) *SearchCrossSellProductsHandler {
	return &SearchCrossSellProductsHandler{
		name:                           name,
		searchCrossSellProductsCommand: command,
	}
}

func (h *SearchCrossSellProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	var (
		request *SearchCrossSellProductsRequest
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

	response, err := h.searchCrossSellProductsCommand.SearchCrossSellProducts(ctx, request)
	if err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	http2.GetSuccessResponseWithBody(w, response)
}

func (_ *SearchCrossSellProductsHandler) getRequestData(r *http.Request) (*SearchCrossSellProductsRequest, error) {
	request := &SearchCrossSellProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
