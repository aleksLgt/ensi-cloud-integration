package recommendationQueryProducts

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
	searchRecommendationQueryProductsCommand interface {
		SearchRecommendationQueryProducts(ctx context.Context, request *SearchRecommendationQueryProductsRequest) (*ensiCloud.SearchRecommendationQueryProductsResponse, error)
	}

	SearchRecommendationQueryProductsHandler struct {
		name                                     string
		searchRecommendationQueryProductsCommand searchRecommendationQueryProductsCommand
	}

	filter struct {
		Query string `json:"query" validate:"nonzero,nonnil"`
	}

	pagination struct {
		Limit int `json:"limit,omitempty" validate:"max=50"`
	}

	SearchRecommendationQueryProductsRequest struct {
		Filter     filter     `json:"filter" validate:"nonnil"`
		Pagination pagination `json:"pagination,omitempty"`
	}
)

func NewSearchRecommendationQueryProductsHandler(command searchRecommendationQueryProductsCommand, name string) *SearchRecommendationQueryProductsHandler {
	return &SearchRecommendationQueryProductsHandler{
		name:                                     name,
		searchRecommendationQueryProductsCommand: command,
	}
}

func (h *SearchRecommendationQueryProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	var (
		request *SearchRecommendationQueryProductsRequest
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

	response, err := h.searchRecommendationQueryProductsCommand.SearchRecommendationQueryProducts(ctx, request)
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

func (_ *SearchRecommendationQueryProductsHandler) getRequestData(r *http.Request) (*SearchRecommendationQueryProductsRequest, error) {
	request := &SearchRecommendationQueryProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
