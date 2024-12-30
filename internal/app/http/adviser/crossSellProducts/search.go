package crossSellProducts

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/validator.v2"

	http2 "ensi-cloud-integration/internal/app/http"
	"ensi-cloud-integration/internal/domain/crossSellProductsDomain"
)

type (
	searchCrossSellProductsCommand interface {
		SearchCrossSellProducts(ctx context.Context, request *crossSellProductsDomain.SearchCrossSellProductsRequest) (*crossSellProductsDomain.SearchCrossSellProductsResponse, error)
	}

	SearchCrossSellProductsHandler struct {
		name                           string
		searchCrossSellProductsCommand searchCrossSellProductsCommand
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
		request *crossSellProductsDomain.SearchCrossSellProductsRequest
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

	searchResponse, err := h.searchCrossSellProductsCommand.SearchCrossSellProducts(ctx, request)
	if err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	buf, err := json.Marshal(&searchResponse)
	if err != nil {
		http2.GetErrorResponse(w, h.name, fmt.Errorf("failed to encode response %w", err), http.StatusInternalServerError)
	}

	http2.GetSuccessResponseWithBody(w, buf)
}

func (_ *SearchCrossSellProductsHandler) getRequestData(r *http.Request) (*crossSellProductsDomain.SearchCrossSellProductsRequest, error) {
	request := &crossSellProductsDomain.SearchCrossSellProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
