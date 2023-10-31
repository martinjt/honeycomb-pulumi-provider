package provider

import (
	"fmt"

	hnyclient "github.com/honeycombio/terraform-provider-honeycombio/client"
	p "github.com/pulumi/pulumi-go-provider"
)

type HoneycombProviderConfig struct {
	Client *hnyclient.Client

	Domain string `pulumi:"domain"`
	ApiKey string `pulumi:"apikey"`
}

func (config *HoneycombProviderConfig) Configure(ctx p.Context) error {
	client, _ := hnyclient.NewClient(&hnyclient.Config{
		APIKey: config.ApiKey,
		APIUrl: fmt.Sprintf("https://%v", config.Domain),
	})
	config.Client = client
	return nil
}
