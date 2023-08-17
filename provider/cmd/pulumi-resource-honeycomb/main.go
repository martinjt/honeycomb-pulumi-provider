// Copyright 2016-2022, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"encoding/json"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"io"
	"net/http"
	"time"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

func main() {
	err := p.RunProvider("honeycomb", Version,
		// We tell the provider what resources it needs to support.
		// In this case, a single custom resource.
		infer.Provider(infer.Options{
			Resources: []infer.InferredResource{
				infer.Resource[Dataset, DatasetArgs, DatasetState](),
			},
			Config: infer.Config[*Config](),
		}))
	if err != nil {
		return
	}
}

type Config struct {
	ApiKey string `pulumi:"apikey"`
}

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type Dataset struct{}

// Each resource has in input struct, defining what arguments it accepts.
type DatasetArgs struct {
	// Fields projected into Pulumi must be public and hava a `pulumi:"..."` tag.
	// The pulumi tag doesn't need to match the field name, but its generally a
	// good idea.
	Name            string `pulumi:"name"`
	Description     string `pulumi:"description"`
	ExpandJsonDepth int    `pulumi:"expand_json_depth"`
}

// Each resource has a state, describing the fields that exist on the created resource.
type DatasetState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	DatasetArgs
	// Here we define a required output called result.
	Name            string `pulumi:"name"`
	Description     string `pulumi:"description"`
	ExpandJsonDepth int    `pulumi:"expand_json_depth"`
	Slug            string `pulumi:"slug"`
	CreatedAt       string `pulumi:"createdAd"`
}

type DatasetCreateRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ExpandJsonDepth int    `json:"expand_json_depth"`
}

// All resources must implement Create at a minumum.
func (Dataset) Create(ctx p.Context, name string, input DatasetArgs, preview bool) (string, DatasetState, error) {
	config := infer.GetConfig[Config](ctx)
	state := DatasetState{DatasetArgs: input}
	if preview {
		return name, state, nil
	}
	makeDataset(input, config)
	return name, state, nil
}

func makeDataset(args DatasetArgs, config Config) {

	datasetRequest := DatasetCreateRequest{
		Name:            args.Name,
		Description:     args.Description,
		ExpandJsonDepth: args.ExpandJsonDepth,
	}

	marshalledRequest, err := json.Marshal(datasetRequest)
	req, err := http.NewRequest(
		"http://api.honeycomb.io/v1/datasets",
		"application/x-www-form-urlencoded",
		bytes.NewReader(marshalledRequest))
	req.Header.Set("X-Honeycomb-Team", config.ApiKey)

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		// handle error
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
}
