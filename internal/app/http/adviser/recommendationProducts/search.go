package recommendationProducts

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
	searchRecommendationProductsCommand interface {
		SearchRecommendationProducts(ctx context.Context, request *SearchRecommendationProductsRequest) (*ensiCloud.SearchRecommendationProductsResponse, error)
	}

	SearchRecommendationProductsHandler struct {
		name                                string
		searchRecommendationProductsCommand searchRecommendationProductsCommand
	}

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
)

func NewSearchRecommendationProductsHandler(command searchRecommendationProductsCommand, name string) *SearchRecommendationProductsHandler {
	return &SearchRecommendationProductsHandler{
		name:                                name,
		searchRecommendationProductsCommand: command,
	}
}

func (h *SearchRecommendationProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	var (
		request *SearchRecommendationProductsRequest
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

	response, err := h.searchRecommendationProductsCommand.SearchRecommendationProducts(ctx, request)
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

func (_ *SearchRecommendationProductsHandler) getRequestData(r *http.Request) (*SearchRecommendationProductsRequest, error) {
	request := &SearchRecommendationProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
