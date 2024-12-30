package ensiCloud

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"ensi-cloud-integration/internal/app/http/catalog"
)

type (
	SearchCatalogResponse struct {
		Data struct {
			TotalProducts   int      `json:"total_products"`
			Products        []string `json:"products"`
			TotalCategories int      `json:"total_categories"`
			Categories      []string `json:"categories"`
			Correction      string   `json:"correction"`
			ProductHints    []struct {
				Word string `json:"word"`
				Hint string `json:"hint"`
			} `json:"product_hints"`
			Filters []struct {
				Name   string `json:"name"`
				Code   string `json:"code"`
				Values []struct {
					Id   string `json:"id"`
					Name string `json:"name"`
				} `json:"values"`
			} `json:"filters"`
		} `json:"data"`
	}
)

const SearchCatalogPath = "/api/v1/catalog/search"

func (c *Client) SearchCatalog(ctx context.Context, request *catalog.SearchCatalogRequest) (*SearchCatalogResponse, error) {
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

	response := &SearchCatalogResponse{}
	err = json.NewDecoder(httpResponse.Body).Decode(response)
	if err != nil {
		return nil, fmt.Errorf("failed to decode error response: %w", err)
	}

	return response, nil
}
