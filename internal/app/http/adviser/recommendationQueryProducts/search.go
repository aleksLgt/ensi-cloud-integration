package recommendationQueryProducts

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/validator.v2"

	http2 "ensi-cloud-integration/internal/app/http"
	"ensi-cloud-integration/internal/domain"
)

type (
	searchRecommendationQueryProductsCommand interface {
		SearchRecommendationQueryProducts(
			ctx context.Context,
			request *domain.SearchRecommendationQueryProductsRequest,
		) (*domain.SearchRecommendationQueryProductsResponse, error)
	}

	SearchRecommendationQueryProductsHandler struct {
		name                                     string
		searchRecommendationQueryProductsCommand searchRecommendationQueryProductsCommand
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
		request *domain.SearchRecommendationQueryProductsRequest
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

	searchResponse, err := h.searchRecommendationQueryProductsCommand.SearchRecommendationQueryProducts(ctx, request)
	if err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	buf, err := json.Marshal(&searchResponse)
	if err != nil {
		http2.GetErrorResponse(w, h.name, fmt.Errorf("failed to encode response %w", err), http.StatusInternalServerError)
	}

	if len(searchResponse.Errors) > 0 {
		http2.GetResponseWithBody(w, buf, http.StatusBadRequest)
		return
	}

	http2.GetResponseWithBody(w, buf, http.StatusOK)
}

func (_ *SearchRecommendationQueryProductsHandler) getRequestData(r *http.Request) (*domain.SearchRecommendationQueryProductsRequest, error) {
	request := &domain.SearchRecommendationQueryProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
