package honeycombapi

type HoneycombApiConfig struct {
	Domain string `pulumi:"domain"`
	ApiKey string `pulumi:"apikey"`
}
