package catalog

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/validator.v2"

	http2 "ensi-cloud-integration/internal/app/http"
	"ensi-cloud-integration/internal/app/http/rules"
	"ensi-cloud-integration/internal/domain"
)

type (
	searchCatalogCommand interface {
		SearchCatalog(ctx context.Context, request *domain.SearchCatalogRequest) (*domain.SearchCatalogResponse, error)
	}

	SearchCatalogHandler struct {
		name                 string
		searchCatalogCommand searchCatalogCommand
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
		request *domain.SearchCatalogRequest
		err     error
	)

	if request, err = h.getRequestData(r); err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err = validator.SetValidationFunc("sortTypeRule", rules.SortTypeRule); err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err = validator.Validate(request); err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	searchResponse, err := h.searchCatalogCommand.SearchCatalog(ctx, request)
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

func (_ *SearchCatalogHandler) getRequestData(r *http.Request) (*domain.SearchCatalogRequest, error) {
	request := &domain.SearchCatalogRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
