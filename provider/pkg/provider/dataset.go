package provider

import (
	"fmt"

	hnyclient "github.com/honeycombio/terraform-provider-honeycombio/client"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

type Dataset struct{}

type DatasetInputs struct {
	Name            string `pulumi:"name"`
	Description     string `pulumi:"description"`
	ExpandJsonDepth int    `pulumi:"expand_json_depth"`
}

type DatasetState struct {
	DatasetInputs

	Name            string `pulumi:"name"`
	Description     string `pulumi:"description"`
	ExpandJsonDepth int    `pulumi:"expand_json_depth"`
	Slug            string `pulumi:"slug"`
	CreatedAt       string `pulumi:"createdAt"`
}

func (Dataset) Create(ctx p.Context, name string, input DatasetInputs, preview bool) (string, DatasetState, error) {
	honeycombProviderConfig := infer.GetConfig[HoneycombProviderConfig](ctx)
	honeycombProviderConfig.Configure(ctx)

	state := DatasetState{DatasetInputs: input}
	if preview {
		return name, state, nil
	}

	apiClient := honeycombProviderConfig.Client

	datasetResponse, _ := apiClient.Datasets.Create(ctx, &hnyclient.Dataset{
		Name:            input.Name,
		Description:     input.Description,
		ExpandJSONDepth: input.ExpandJsonDepth,
	})

	state.Slug = datasetResponse.Slug
	state.Name = datasetResponse.Name
	state.Description = datasetResponse.Description
	state.ExpandJsonDepth = datasetResponse.ExpandJSONDepth

	ctx.Log(diag.Warning, fmt.Sprintf("Created dataset %v with slug \"%v\"", input.Name, state.Slug))
	return name, state, nil
}
