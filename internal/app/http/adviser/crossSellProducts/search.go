package crossSellProducts

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/validator.v2"

	http2 "ensi-cloud-integration/internal/app/http"
	"ensi-cloud-integration/internal/clients/ensiCloud"
)

type (
	searchCrossSellProductsCommand interface {
		SearchCrossSellProducts(ctx context.Context, request *SearchCrossSellProductsRequest) (*ensiCloud.SearchCrossSellProductsResponse, error)
	}

	SearchCrossSellProductsHandler struct {
		name                           string
		searchCrossSellProductsCommand searchCrossSellProductsCommand
	}

	filter struct {
		ProductId string `json:"product_id" validate:"nonzero,nonnil"`
	}

	pagination struct {
		Limit int `json:"limit,omitempty" validate:"max=50"`
	}

	SearchCrossSellProductsRequest struct {
		Filter     filter     `json:"filter" validate:"nonnil"`
		Pagination pagination `json:"pagination,omitempty"`
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

	buf, err := json.Marshal(&response)
	if err != nil {
		http2.GetErrorResponse(w, h.name, fmt.Errorf("failed to encode response %w", err), http.StatusInternalServerError)
	}

	http2.GetSuccessResponseWithBody(w, buf)
}

func (_ *SearchCrossSellProductsHandler) getRequestData(r *http.Request) (*SearchCrossSellProductsRequest, error) {
	request := &SearchCrossSellProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
