package provider

import (
	"fmt"

	hnyclient "github.com/honeycombio/terraform-provider-honeycombio/client"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

type Column struct{}

type ColumnInputs struct {
	Dataset string `pulumi:"dataset"`

	Name        string     `pulumi:"name"`
	Type        ColumnType `pulumi:"type"`
	Description *string    `pulumi:"description,optional"`
	Hidden      *bool      `pulumi:"hidden,optional"`
}

type ColumnState struct {
	ColumnInputs

	Id          string                `pulumi:"columnId"`
	Name        string                `pulumi:"name"`
	Type        *hnyclient.ColumnType `pulumi:"type"`
	Description *string               `pulumi:"description"`
	Hidden      *bool                 `pulumi:"hidden"`
}

type ColumnType string

var _ = (infer.Enum[ColumnType])((*ColumnType)(nil))

const (
	String  ColumnType = "string"
	Integer ColumnType = "integer"
	Float   ColumnType = "float"
	Boolean ColumnType = "boolean"
)

func (*ColumnType) Values() []infer.EnumValue[ColumnType] {
	return []infer.EnumValue[ColumnType]{
		{Value: String, Description: "String Columns are the catch all"},
		{Value: Integer, Description: "Int Columns are for integers"},
		{Value: Float, Description: "Float Columns are for floats"},
		{Value: Boolean, Description: "Boolean Columns are for booleans"},
	}
}

func (columnType ColumnType) convertToHoneycombColumnType() *hnyclient.ColumnType {
	switch columnType {
	case String:
		return hnyclient.ToPtr(hnyclient.ColumnTypeString)
	case Integer:
		return hnyclient.ToPtr(hnyclient.ColumnTypeInteger)
	case Float:
		return hnyclient.ToPtr(hnyclient.ColumnTypeFloat)
	case Boolean:
		return hnyclient.ToPtr(hnyclient.ColumnTypeBoolean)
	default:
		return hnyclient.ToPtr(hnyclient.ColumnTypeString)
	}
}

func (*Column) Create(ctx p.Context, name string, input ColumnInputs, preview bool) (string, ColumnState, error) {
	honeycombProviderConfig := infer.GetConfig[HoneycombProviderConfig](ctx)
	honeycombProviderConfig.Configure(ctx)

	state := ColumnState{ColumnInputs: input}
	if preview {
		return name, state, nil
	}

	apiClient := honeycombProviderConfig.Client
	ctx.Log(diag.Warning, fmt.Sprintf("Creating column %v on %v", input.Name, input.Dataset))
	columnResponse, err := apiClient.Columns.Create(ctx, input.Dataset, &hnyclient.Column{
		KeyName:     input.Name,
		Type:        input.Type.convertToHoneycombColumnType(),
		Description: *input.Description,
		Hidden:      input.Hidden,
	})

	state.Id = columnResponse.ID
	state.Name = columnResponse.KeyName
	state.Type = columnResponse.Type
	state.Description = &columnResponse.Description
	state.Hidden = columnResponse.Hidden

	if err != nil {
		return "", state, err
	}

	return columnResponse.ID, state, nil
}

func (*Column) Read(ctx p.Context, id string, input ColumnInputs) (ColumnState, error) {
	honeycombProviderConfig := infer.GetConfig[HoneycombProviderConfig](ctx)
	honeycombProviderConfig.Configure(ctx)

	apiClient := honeycombProviderConfig.Client

	columnResponse, err := apiClient.Columns.Get(ctx, input.Dataset, id)

	if err != nil {
		return ColumnState{}, err
	}

	return ColumnState{
		Id:          columnResponse.ID,
		Name:        columnResponse.KeyName,
		Type:        columnResponse.Type,
		Description: &columnResponse.Description,
		Hidden:      columnResponse.Hidden,
	}, nil
}

func (*Column) Diff(ctx p.Context, id string, olds ColumnState, news ColumnInputs) (p.DiffResponse, error) {
	ctx.Log(diag.Warning, fmt.Sprintf("Diffing column %v on %v", id, news.Dataset))
	diff := map[string]p.PropertyDiff{}

	if *olds.Type != *news.Type.convertToHoneycombColumnType() {
		ctx.Log(diag.Warning, fmt.Sprintf("Type changed from %v to %v", *olds.Type, *news.Type.convertToHoneycombColumnType()))
		diff["type"] = p.PropertyDiff{Kind: p.Update}
	}
	if *olds.Description != *news.Description {
		diff["description"] = p.PropertyDiff{Kind: p.Update}
	}
	if olds.Hidden != news.Hidden {
		diff["hidden"] = p.PropertyDiff{Kind: p.Update}
	}
	if olds.Name != news.Name {
		diff["name"] = p.PropertyDiff{Kind: p.UpdateReplace}
	}

	return p.DiffResponse{
		DeleteBeforeReplace: true,
		HasChanges:          len(diff) > 0,
		DetailedDiff:        diff,
	}, nil
}

func (*Column) Update(ctx p.Context, id string, currentState ColumnState, input ColumnInputs, preview bool) (ColumnState, error) {

	if preview {
		return currentState, nil
	}

	honeycombProviderConfig := infer.GetConfig[HoneycombProviderConfig](ctx)
	honeycombProviderConfig.Configure(ctx)

	apiClient := honeycombProviderConfig.Client

	columnResponse, err := apiClient.Columns.Update(ctx, input.Dataset, &hnyclient.Column{
		ID:          id,
		KeyName:     input.Name,
		Description: *input.Description,
		Hidden:      input.Hidden,
	})

	if err != nil {
		return currentState, err
	}

	currentState.Name = columnResponse.KeyName
	currentState.Type = columnResponse.Type
	currentState.Description = &columnResponse.Description
	currentState.Hidden = columnResponse.Hidden

	return currentState, nil
}

func (*Column) Delete(ctx p.Context, id string, input ColumnState) error {
	honeycombProviderConfig := infer.GetConfig[HoneycombProviderConfig](ctx)
	honeycombProviderConfig.Configure(ctx)

	apiClient := honeycombProviderConfig.Client

	ctx.Log(diag.Warning, fmt.Sprintf("Deleting column %v on %v", id, input.Dataset))

	err := apiClient.Columns.Delete(ctx, input.Dataset, id)

	return err
}
