package ensiCloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"ensi-cloud-integration/internal/app/http/adviser/recommendationProducts"
)

const SearchRecommendationProductsPath = "/api/v1/adviser/recommendation-products:search"

func (c *Client) SearchRecommendationProducts(
	ctx context.Context,
	request *recommendationProducts.SearchRecommendationProductsRequest,
) ([]byte, error) {
	data, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request %w", err)
	}

	path, err := url.JoinPath(c.basePath, SearchRecommendationProductsPath)
	if err != nil {
		return nil, fmt.Errorf("incorrect base basePath: %w", err)
	}

	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	httpRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.privateToken))

	client := &http.Client{}

	httpResponse, err := client.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %w", err)
	}

	log.Println(httpRequest, httpResponse)

	defer func() {
		_ = httpResponse.Body.Close()
	}()

	if httpResponse.StatusCode != http.StatusOK {
		response := &ErrorResponse{}
		err = json.NewDecoder(httpResponse.Body).Decode(response)

		if err != nil {
			return nil, fmt.Errorf("failed to decode error response: %w", err)
		}

		return nil, fmt.Errorf("HTTP request responded with: %d , message: %s", httpResponse.StatusCode, response)
	}

	// TODO response

	return nil, nil
}
