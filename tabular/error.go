package tabular

import (
	"encoding/json"
	"io"
)

type JSONErrorResponse struct {
	Error JSONError `json:"error"`
}

type JSONError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    int    `json:"code"`
}

func ParseErrorResponse(response io.ReadCloser) (JSONErrorResponse, error) {
	var errResp JSONErrorResponse
	body, err := io.ReadAll(response)
	if err != nil {
		return JSONErrorResponse{}, err
	}
	err = json.Unmarshal(body, &errResp)
	if err != nil {
		return JSONErrorResponse{}, err
	}

	return errResp, nil
}
