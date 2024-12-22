package ensiCloud

import (
	"context"
	"errors"
)

type (
	Client struct {
		privateToken, publicToken, basePath string
	}
)

func New(basePath, privateToken, publicToken string) (*Client, error) {
	if privateToken == "" {
		return nil, errors.New("ensi cloud service has empty private token")
	}

	if publicToken == "" {
		return nil, errors.New("ensi cloud service has empty public token")
	}

	return &Client{
		privateToken: privateToken,
		publicToken:  publicToken,
		basePath:     basePath,
	}, nil
}

func (c *Client) IndexProducts(ctx context.Context) error {
	return nil
}
