package http

import (
	"bytes"
	"net/http"
)

func GetErrorResponse(w http.ResponseWriter, handlerName string, err error, statusCode int) {
	w.WriteHeader(statusCode)

	buf := bytes.NewBufferString(handlerName)
	buf.WriteString(": ")
	buf.WriteString(err.Error())
	buf.WriteString("\n")
	_, _ = w.Write(buf.Bytes())
}

func GetSuccessResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func GetNoContentResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func GetResponseWithBody(w http.ResponseWriter, body []byte, status int) {
	w.WriteHeader(status)

	_, _ = w.Write(body)
}
