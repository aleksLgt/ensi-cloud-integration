package ensiCloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"ensi-cloud-integration/internal/app/http/indexes/products"
)

const IndexProductsPath = "/api/v1/indexes/products"

func (c *Client) IndexProducts(ctx context.Context, request *products.IndexProductsRequest) error {
	data, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to encode request %w", err)
	}

	path, err := url.JoinPath(c.basePath, IndexProductsPath)
	if err != nil {
		return fmt.Errorf("incorrect base basePath: %w", err)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.privateToken))

	client := &http.Client{}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	log.Println(httpRequest, httpResponse)

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
