package ensiCloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"ensi-cloud-integration/internal/app/http/indexes/categories"
)

const IndexCategoriesPath = "/api/v1/indexes/categories"

func (c *Client) IndexCategories(ctx context.Context, request *categories.IndexCategoriesRequest) error {
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to encode request %w", err)
	}

	path, err := url.JoinPath(c.basePath, IndexCategoriesPath)
	if err != nil {
		return fmt.Errorf("incorrect base basePath: %w", err)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.privateToken))
	httpRequest.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	defer func() {
		_ = httpResponse.Body.Close()
	}()

	if httpResponse.StatusCode != http.StatusOK {
		response := &ErrorResponse{}
		err = json.NewDecoder(httpResponse.Body).Decode(response)

		if err != nil {
			return fmt.Errorf("failed to decode error response: %w", err)
		}

		return fmt.Errorf("HTTP request responded with: %d , message: %s", httpResponse.StatusCode, response)
	}

	return nil
}
