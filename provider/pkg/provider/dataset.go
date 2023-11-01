package provider

import (
	hnyclient "github.com/honeycombio/terraform-provider-honeycombio/client"
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
	honeycombProviderConfig := infer.GetConfig[HoneycombProviderConfig](ctx)
	state := DatasetState{DatasetArgs: input}
	if preview {
		return name, state, nil
	}
	makeDataset(ctx, honeycombProviderConfig.Client, input, state)
	return name, state, nil
}

func (Dataset) Annotate(a infer.Annotator) {
	a.SetToken("Resources", "Dataset")
}

func makeDataset(ctx p.Context, apiClient *hnyclient.Client, args DatasetArgs, state DatasetState) {

	datasetResponse, _ := apiClient.Datasets.Create(ctx, &hnyclient.Dataset{
		Name:            args.Name,
		Description:     args.Description,
		ExpandJSONDepth: args.ExpandJsonDepth,
	})

	state.Slug = datasetResponse.Slug
}
