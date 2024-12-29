package products

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/validator.v2"

	http2 "ensi-cloud-integration/internal/app/http"
)

type (
	indexProductsCommand interface {
		IndexProducts(ctx context.Context, request *IndexProductsRequest) error
	}

	IndexProductsHandler struct {
		name                 string
		indexProductsCommand indexProductsCommand
	}

	propertyRequest struct {
		Name   string   `json:"name" validate:"nonzero,nonnil"`
		Values []string `json:"values" validate:"nonnil,min=1"`
	}

	locationRequest struct {
		Id    string `json:"id" validate:"nonzero,nonnil"`
		Price int    `json:"price" validate:"nonzero,nonnil"`
	}

	bodyRequest struct {
		Name        string            `json:"name" validate:"nonzero,nonnil"`
		URL         string            `json:"url"`
		CategoryIds []string          `json:"category_ids" validate:"nonnil,min=1"`
		Brand       string            `json:"brand"`
		VendorCode  string            `json:"vendor_code" validate:"nonzero,nonnil"`
		Barcodes    []string          `json:"barcodes"`
		Description string            `json:"description"`
		Picture     string            `json:"picture"`
		Country     string            `json:"country"`
		GroupIds    []string          `json:"group_ids"`
		Locations   []locationRequest `json:"locations" validate:"nonnil"`
		Properties  []propertyRequest `json:"properties" validate:"nonnil"`
	}

	actionRequest struct {
		Action string      `json:"action"` // TODO custom rule
		Id     string      `json:"id" validate:"nonzero,nonnil"`
		Body   bodyRequest `json:"body" validate:"nonnil"`
	}

	IndexProductsRequest struct {
		Actions []actionRequest `json:"actions" validate:"nonnil"`
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
		request *IndexProductsRequest
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

	err = h.indexProductsCommand.IndexProducts(ctx, request)
	if err != nil {
		http2.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}

	http2.GetSuccessResponse(w)
}

func (_ *IndexProductsHandler) getRequestData(r *http.Request) (*IndexProductsRequest, error) {
	request := &IndexProductsRequest{}
	err := json.NewDecoder(r.Body).Decode(request)

	if err != nil {
		return nil, fmt.Errorf("failed to decode request data: %w", err)
	}

	return request, nil
}
