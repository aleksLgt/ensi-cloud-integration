package products

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
	indexProductsCommand interface {
		IndexProducts(ctx context.Context, request *domain.IndexProductsRequest) (*domain.IndexProductsResponse, error)
	}

	IndexProductsHandler struct {
		name                 string
		indexProductsCommand indexProductsCommand
	}
)

func NewIndexProductsHandler(command indexProductsCommand, name string) *IndexProductsHandler {
	return &IndexProductsHandler{
		name:                 name,
		indexProductsCommand: command,
	}
}

func (h *IndexProductsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	ctx := r.Context()

	var (
		request *domain.IndexProductsRequest
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

	response, err := h.indexProductsCommand.IndexProducts(ctx, request)
	if err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	buf, err := json.Marshal(&response)
	if err != nil {
		http2.GetErrorResponse(w, h.name, fmt.Errorf("failed to encode response %w", err), http.StatusInternalServerError)
	}

	if len(response.Errors) > 0 {
		http2.GetResponseWithBody(w, buf, http.StatusBadRequest)
		return
	}

	http2.GetResponseWithBody(w, buf, http.StatusOK)
}

func (_ *IndexProductsHandler) getRequestData(r *http.Request) (*domain.IndexProductsRequest, error) {
	request := &domain.IndexProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
