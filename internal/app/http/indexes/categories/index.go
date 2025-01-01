package categories

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
	indexCategoriesCommand interface {
		IndexCategories(ctx context.Context, request *domain.IndexCategoriesRequest) (*domain.IndexCategoriesResponse, error)
	}

	IndexCategoriesHandler struct {
		name                   string
		indexCategoriesCommand indexCategoriesCommand
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
		request *domain.IndexCategoriesRequest
		err     error
	)

	if request, err = h.getRequestData(r); err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err = validator.SetValidationFunc("actionTypeRule", rules.ActionTypeRule); err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	if err = validator.Validate(request); err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	response, err := h.indexCategoriesCommand.IndexCategories(ctx, request)
	if err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	buf, err := json.Marshal(&response)
	if err != nil {
		http2.GetErrorResponse(w, h.name, fmt.Errorf("failed to encode response %w", err), http.StatusInternalServerError)
		return
	}

	if len(response.Errors) > 0 {
		http2.GetResponseWithBody(w, buf, http.StatusBadRequest)
		return
	}

	http2.GetResponseWithBody(w, buf, http.StatusOK)
}

func (_ *IndexCategoriesHandler) getRequestData(r *http.Request) (*domain.IndexCategoriesRequest, error) {
	request := &domain.IndexCategoriesRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
