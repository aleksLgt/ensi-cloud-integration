package recommendationProducts

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/validator.v2"

	http2 "ensi-cloud-integration/internal/app/http"
	"ensi-cloud-integration/internal/domain/recommendationProductsDomain"
)

type (
	searchRecommendationProductsCommand interface {
		SearchRecommendationProducts(ctx context.Context, request *recommendationProductsDomain.SearchRecommendationProductsRequest) (*recommendationProductsDomain.SearchRecommendationProductsResponse, error)
	}

	SearchRecommendationProductsHandler struct {
		name                                string
		searchRecommendationProductsCommand searchRecommendationProductsCommand
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
		request *recommendationProductsDomain.SearchRecommendationProductsRequest
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

	searchResponse, err := h.searchRecommendationProductsCommand.SearchRecommendationProducts(ctx, request)
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

func (_ *SearchRecommendationProductsHandler) getRequestData(r *http.Request) (*recommendationProductsDomain.SearchRecommendationProductsRequest, error) {
	request := &recommendationProductsDomain.SearchRecommendationProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
