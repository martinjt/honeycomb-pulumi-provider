package honeycombapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	HEADER_CONTENT_TYPE = "Content-Type"
	HONEYCOMB_API_KEY   = "X-Honeycomb-Team"
)

type HoneycombApi struct {
	HttpClient http.Client
	Config     HoneycombApiConfig
}

func (api HoneycombApi) Setup(config HoneycombApiConfig) {

}

func SendPostRequest[A any](api *HoneycombApi, route string, body any) (A, ApiError) {
	requestJsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST",
		api.routeUrl(route),
		bytes.NewReader(requestJsonBody))

	req.Header.Set(HONEYCOMB_API_KEY, "application/json; charset=UTF-8")
	req.Header.Set(HEADER_CONTENT_TYPE, api.Config.ApiKey)

	response, _ := api.HttpClient.Do(req)
	defer response.Body.Close()

	parsedResponse, parseError := parseResponse[A](response)

	return parsedResponse, parseError
}

func (api *HoneycombApi) routeUrl(route string) string {
	const baseUrl = "https://%v/1/%v"
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
