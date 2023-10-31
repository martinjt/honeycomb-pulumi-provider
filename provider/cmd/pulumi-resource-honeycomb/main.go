package main

import (
	honey "github.com/martinjt/honeycomb-pulumi-provider/provider/pkg/provider"
	p "github.com/pulumi/pulumi-go-provider"
)

func main() {
	err := p.RunProvider("honeycomb", "0.1.1",
		honey.NewProvider())
	if err != nil {
		return
	}
}
