package honeycombapi

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Dataset struct{}

type DatasetArgs struct {
	Name            string `pulumi:"name"`
	Description     string `pulumi:"description"`
	ExpandJsonDepth int    `pulumi:"expand_json_depth"`
}

type DatasetState struct {
	DatasetArgs

	Name            string `pulumi:"name"`
	Description     string `pulumi:"description"`
	ExpandJsonDepth int    `pulumi:"expand_json_depth"`
	Slug            string `pulumi:"slug"`
	CreatedAt       string `pulumi:"createdAt"`
}

func (Dataset) Create(ctx p.Context, name string, input DatasetArgs, preview bool) (string, DatasetState, error) {
	honeycombApiConfig = infer.GetConfig[HoneycombApiConfig](ctx)
	state := DatasetState{DatasetArgs: input}
	if preview {
		return name, state, nil
	}
	makeDataset(input, state)
	return name, state, nil
}

func makeDataset(args DatasetArgs, state DatasetState) {

	datasetRequest := DatasetCreateRequest{
		Name:        args.Name,
		Description: args.Description,
	}
	datasetCreateResponse, _ := sendPostRequestToHoneycomb[DatasetCreateResponse]("datasets", datasetRequest)

	state.Slug = datasetCreateResponse.Slug
}

type DatasetCreateRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ExpandJsonDepth int    `json:"expand_json_depth"`
}

type DatasetCreateResponse struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ExpandJsonDepth int    `json:"expand_json_depth"`
	Slug            string `json:"slug"`
}
