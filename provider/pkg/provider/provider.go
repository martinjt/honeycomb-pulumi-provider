package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
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
		Config: infer.Config[HoneycombProviderConfig](),
		Resources: []infer.InferredResource{
			infer.Resource[Dataset](),
		},
	})
}
