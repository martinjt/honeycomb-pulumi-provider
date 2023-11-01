package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

func NewProvider() p.Provider {
	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			DisplayName: "Honeycomb",
			Description: "The Honeycomb Pulumi provider makes it simple to create honeycomb resources during your pipeline",
			Keywords: []string{
				"honeycomb",
				"observability",
				"category/observability",
				"kind/native",
			},
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "Resources",
		},
		Config: infer.Config[*HoneycombProviderConfig](),
		Resources: []infer.InferredResource{
			infer.Resource[Dataset, DatasetInputs, DatasetState](),
			infer.Resource[*Column, ColumnInputs, ColumnState](),
		},
	})
}
