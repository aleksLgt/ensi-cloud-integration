package ensiCloud

import (
	"errors"
)

type (
	Client struct {
		privateToken, publicToken, basePath string
	}

	dataResponse struct{}

	metaResponse struct{}

	errorErrorResponse struct {
		Code    string       `json:"code"`
		Message string       `json:"message"`
		Meta    metaResponse `json:"meta"`
	}

	ErrorResponse struct {
		Data   dataResponse         `json:"data,omitempty"`
		Errors []errorErrorResponse `json:"errors,omitempty"`
		Meta   metaResponse         `json:"meta,omitempty"`
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
