package provider

import (
	"fmt"

	hnyclient "github.com/honeycombio/terraform-provider-honeycombio/client"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

type HoneycombProviderConfig struct {
	Client *hnyclient.Client

	ApiKey string `pulumi:"apikey"`

	Domain  *string `pulumi:"domain,optional"`
	Version *string `pulumi:"version,optional"`
}

func (config *HoneycombProviderConfig) Configure(ctx p.Context) error {
	ctx.Log(diag.Info, "Configuring Honeycomb provider")
	clientConfig := &hnyclient.Config{
		APIKey: config.ApiKey,
	}
	if config.Domain != nil {
		clientConfig.APIUrl = fmt.Sprintf("https://%v", config.Domain)
	}
	client, _ := hnyclient.NewClient(clientConfig)
	config.Client = client
	return nil
}
