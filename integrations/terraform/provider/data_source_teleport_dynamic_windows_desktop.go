// Code generated by _gen/main.go DO NOT EDIT
/*
Copyright 2015-2024 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package provider

import (
	"context"

	apitypes "github.com/gravitational/teleport/api/types"
	"github.com/gravitational/trace"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/gravitational/teleport/integrations/terraform/tfschema"
)

// dataSourceTeleportDynamicWindowsDesktopType is the data source metadata type
type dataSourceTeleportDynamicWindowsDesktopType struct{}

// dataSourceTeleportDynamicWindowsDesktop is the resource
type dataSourceTeleportDynamicWindowsDesktop struct {
	p Provider
}

// GetSchema returns the data source schema
func (r dataSourceTeleportDynamicWindowsDesktopType) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfschema.GenSchemaDynamicWindowsDesktopV1(ctx)
}

// NewDataSource creates the empty data source
func (r dataSourceTeleportDynamicWindowsDesktopType) NewDataSource(_ context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return dataSourceTeleportDynamicWindowsDesktop{
		p: *(p.(*Provider)),
	}, nil
}

// Read reads teleport DynamicWindowsDesktop
func (r dataSourceTeleportDynamicWindowsDesktop) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var id types.String
	diags := req.Config.GetAttribute(ctx, path.Root("metadata").AtName("name"), &id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	desktopI, err := r.p.Client.DynamicDesktopClient().GetDynamicWindowsDesktop(ctx, id.Value)
	if err != nil {
		resp.Diagnostics.Append(diagFromWrappedErr("Error reading DynamicWindowsDesktop", trace.Wrap(err), "dynamic_windows_desktop"))
		return
	}

	var state types.Object
	resp.Diagnostics.Append(req.Config.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Todo: Remove after updating terraform-plugin to >=v1.5.0.
	// terraform-plugin-testing version <1.5.0 requires data resources to
	// implement the 'id' attribute.
	// https://developer.hashicorp.com/terraform/plugin/framework/acctests#no-id-found-in-attributes
	v, ok := state.Attrs["id"]
	if !ok || v.IsNull() {
		state.Attrs["id"] = id
	}

	
	desktop := desktopI.(*apitypes.DynamicWindowsDesktopV1)
	diags = tfschema.CopyDynamicWindowsDesktopV1ToTerraform(ctx, desktop, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
