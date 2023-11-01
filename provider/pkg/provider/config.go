package provider

import (
	"fmt"

	hnyclient "github.com/honeycombio/terraform-provider-honeycombio/client"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

type HoneycombProviderConfig struct {
	Client *hnyclient.Client

	Domain string `pulumi:"domain"`
	ApiKey string `pulumi:"apikey"`
}

func (config *HoneycombProviderConfig) Configure(ctx p.Context) error {
	clientConfig := &hnyclient.Config{
		APIKey: config.ApiKey,
	}
	if config.Domain != "" {
		clientConfig.APIUrl = fmt.Sprintf("https://%v", config.Domain)
	}
	client, _ := hnyclient.NewClient(clientConfig)
	config.Client = client
	return nil
}

func (config *HoneycombProviderConfig) Check(
	ctx p.Context,
	name string,
	oldInputs resource.PropertyMap, newInputs resource.PropertyMap) (HoneycombProviderConfig, []p.CheckFailure, error) {
	return HoneycombProviderConfig{}, []p.CheckFailure{}, nil
}
