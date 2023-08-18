package main

import (
	"github.com/martinjt/honeycomb-pulumi-provider/provider/pkg/version"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

func main() {
	err := p.RunProvider("honeycomb", version.Version,
		infer.Provider(infer.Options{
			Resources: []infer.InferredResource{
				infer.Resource[Dataset, DatasetArgs, DatasetState](),
			},
			ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
				"index": "Dataset",
			},
			Config: infer.Config[*Config](),
		}))
	if err != nil {
		return
	}
}

type Config struct {
	Version string `pulumi:"version,optional"`
	ApiKey  string `pulumi:"apikey"`
}
