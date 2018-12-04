package main

import (
	"encoding/json"
	"fmt"
)

// authenticationInvalid is custom error for Invalid Authentication
type authenticationInvalid struct {
	Code     int
	ErrorStr string
}

func (e *authenticationInvalid) Error() string {
	return fmt.Sprintf("code %d: %s", e.Code, e.ErrorStr)
}

// errorJSONResponse is struct for error messages returned from Twitter server
type errorJSONResponse struct {
	Errors []struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

// CheckForResponseError checks for twitter error response
func CheckForResponseError(response []byte) error {
	var jsonObject errorJSONResponse
	err := json.Unmarshal(response, &jsonObject)
	if err != nil {
		return err
	}
	if len(jsonObject.Errors) > 0 {
		if jsonObject.Errors[0].Code == 32 {
			return &authenticationInvalid{jsonObject.Errors[0].Code, jsonObject.Errors[0].Message}
		}
	}
	return nil
}
