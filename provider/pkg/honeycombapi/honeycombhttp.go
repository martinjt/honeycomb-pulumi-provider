package honeycombapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Headers struct {
	CONTENTTYPE string
}

const (
	HEADER_CONTENT_TYPE = "Content-Type"
	HONEYCOMB_API_KEY   = "X-Honeycomb-Team"
)

var httpClient http.Client
var honeycombApiConfig HoneycombApiConfig

func sendPostRequestToHoneycomb[A any](route string, body any) (A, ApiError) {
	requestJsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST",
		honeycombApiUrl(route),
		bytes.NewReader(requestJsonBody))

	req.Header.Set(HONEYCOMB_API_KEY, "application/json; charset=UTF-8")
	req.Header.Set(HEADER_CONTENT_TYPE, honeycombApiConfig.ApiKey)

	response, _ := httpClient.Do(req)
	defer response.Body.Close()

	parsedResponse, parseError := parseResponse[A](response)

	return parsedResponse, parseError
}

func honeycombApiUrl(route string) string {
	const baseUrl = "https://api.honeycomb.io/1/%v"
	return fmt.Sprintf(baseUrl, route)
}

func parseResponse[A any](httpResponse *http.Response) (apiResponse A, errorResponse ApiError) {
	bodyString, _ := io.ReadAll(httpResponse.Body)
	if parseError := json.Unmarshal(bodyString, apiResponse); parseError != nil {
		errorResponse = ApiError{
			Message: parseError.Error(),
		}
	}
	return
}

type ApiError struct {
	Message string
}
