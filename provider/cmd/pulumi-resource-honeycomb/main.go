package main

import (
	honey "github.com/martinjt/honeycomb-pulumi-provider/provider/pkg/honeycombapi"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

func main() {
	err := p.RunProvider("honeycomb", "0.1.1",
		infer.Provider(infer.Options{
			Resources: []infer.InferredResource{
				infer.Resource[honey.Dataset, honey.DatasetArgs, honey.DatasetState](),
			},
			ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
				"index": "Dataset",
			},
			Config: infer.Config[*honey.Config](),
		}))
	if err != nil {
		return
	}
}
