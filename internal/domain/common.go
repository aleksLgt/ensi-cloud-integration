package domain

type (
	metaResponse struct{}

	errorResponse struct {
		Code    string       `json:"code"`
		Message string       `json:"message"`
		Meta    metaResponse `json:"meta"`
	}
)
