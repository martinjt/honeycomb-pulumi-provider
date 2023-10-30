package honeycombapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Dataset struct{}

// Each resource has in input struct, defining what arguments it accepts.
type DatasetArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but its generally a
	// good idea.
	Name        string `pulumi:"name"`
	Description string `pulumi:"description"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type DatasetState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	DatasetArgs
	// Here we define a required output called result.
	Name            string `pulumi:"name"`
	Description     string `pulumi:"description"`
	ExpandJsonDepth int    `pulumi:"expand_json_depth"`
	Slug            string `pulumi:"slug"`
	CreatedAt       string `pulumi:"createdAd"`
}

func (Dataset) Create(ctx p.Context, name string, input DatasetArgs, preview bool) (string, DatasetState, error) {
	config := infer.GetConfig[Config](ctx)
	state := DatasetState{DatasetArgs: input}
	if preview {
		return name, state, nil
	}
	makeDataset(input, config)
	return name, state, nil
}

func makeDataset(args DatasetArgs, config Config) {

	datasetRequest := DatasetCreateRequest{
		Name:        args.Name,
		Description: args.Description,
	}

	marshalledRequest, _ := json.Marshal(datasetRequest)
	req, _ := http.NewRequest("POST",
		"https://api.honeycomb.io/1/datasets",
		bytes.NewReader(marshalledRequest))
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("X-Honeycomb-Team", config.ApiKey)

	client := &http.Client{}
	response, error := client.Do(req)
	if error != nil {
		panic(error)
	}
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := io.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	defer response.Body.Close()
}

type DatasetCreateRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ExpandJsonDepth int    `json:"expand_json_depth"`
}
