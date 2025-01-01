package ensiCloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"ensi-cloud-integration/internal/domain"
)

const SearchCatalogPath = "/api/v1/catalog/search"

func (c *Client) SearchCatalog(
	ctx context.Context,
	request *domain.SearchCatalogRequest,
) (*domain.SearchCatalogResponse, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request %w", err)
	}

	path, err := url.JoinPath(c.basePath, SearchCatalogPath)
	if err != nil {
		return nil, fmt.Errorf("incorrect base basePath: %w", err)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.publicToken))
	httpRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	defer func() {
		_ = httpResponse.Body.Close()
	}()

	response := &domain.SearchCatalogResponse{}
	err = json.NewDecoder(httpResponse.Body).Decode(response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode error response: %w", err)
	}

	return response, nil
}
