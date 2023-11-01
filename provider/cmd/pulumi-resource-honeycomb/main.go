package main

import (
	honey "github.com/martinjt/honeycomb-pulumi-provider/provider/pkg/provider"
	"github.com/martinjt/honeycomb-pulumi-provider/provider/pkg/version"
	p "github.com/pulumi/pulumi-go-provider"
)

func main() {
	err := p.RunProvider("honeycomb", version.Version,
		honey.NewProvider())
	if err != nil {
		return
	}
}
