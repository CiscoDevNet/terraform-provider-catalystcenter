// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// Section below is generated&owned by "gen/generator.go". //template:begin provider
import (
	"context"
	"os"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	cc "github.com/netascode/go-catalystcenter"
)

// CcProvider defines the provider implementation.
type CcProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// CcProviderModel describes the provider data model.
type CcProviderModel struct {
	Username   types.String `tfsdk:"username"`
	Password   types.String `tfsdk:"password"`
	URL        types.String `tfsdk:"url"`
	Insecure   types.Bool   `tfsdk:"insecure"`
	Retries    types.Int64  `tfsdk:"retries"`
	MaxTimeout types.Int64  `tfsdk:"max_timeout"`
}

// CcProviderData describes the data maintained by the provider.
type CcProviderData struct {
	Client *cc.Client
}

// Metadata returns the provider type name.
func (p *CcProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "catalystcenter"
	resp.Version = p.version
}

func (p *CcProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				MarkdownDescription: "Username for the Catalyst Center instance. This can also be set as the CC_USERNAME environment variable.",
				Optional:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Password for the Catalyst Center instance. This can also be set as the CC_PASSWORD environment variable.",
				Optional:            true,
				Sensitive:           true,
			},
			"url": schema.StringAttribute{
				MarkdownDescription: "URL of the Catalyst Center instance. This can also be set as the CC_URL environment variable.",
				Optional:            true,
			},
			"insecure": schema.BoolAttribute{
				MarkdownDescription: "Allow insecure HTTPS client. This can also be set as the CC_INSECURE environment variable. Defaults to `true`.",
				Optional:            true,
			},
			"retries": schema.Int64Attribute{
				MarkdownDescription: "Number of retries for REST API calls. This can also be set as the CC_RETRIES environment variable. Defaults to `3`.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 9),
				},
			},
			"max_timeout": schema.Int64Attribute{
				MarkdownDescription: "Timeout in seconds for asynchronous tasks. This can also be set as the CC_MAX_TIMEOUT environment variable. Defaults to `60`.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.Between(0, 3600),
				},
			},
		},
	}
}

func (p *CcProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config CcProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// User must provide a username to the provider
	var username string
	if config.Username.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.IsNull() {
		username = os.Getenv("CC_USERNAME")
	} else {
		username = config.Username.ValueString()
	}

	if username == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	// User must provide a password to the provider
	var password string
	if config.Password.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.IsNull() {
		password = os.Getenv("CC_PASSWORD")
	} else {
		password = config.Password.ValueString()
	}

	if password == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find password",
			"Password cannot be an empty string",
		)
		return
	}

	// User must provide a username to the provider
	var url string
	if config.URL.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as url",
		)
		return
	}

	if config.URL.IsNull() {
		url = os.Getenv("CC_URL")
	} else {
		url = config.URL.ValueString()
	}

	if url == "" {
		// Error vs warning - empty value must stop execution
		resp.Diagnostics.AddError(
			"Unable to find url",
			"URL cannot be an empty string",
		)
		return
	}

	var insecure bool
	if config.Insecure.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as insecure",
		)
		return
	}

	if config.Insecure.IsNull() {
		insecureStr := os.Getenv("CC_INSECURE")
		if insecureStr == "" {
			insecure = true
		} else {
			insecure, _ = strconv.ParseBool(insecureStr)
		}
	} else {
		insecure = config.Insecure.ValueBool()
	}

	var retries int64
	if config.Retries.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as retries",
		)
		return
	}

	if config.Retries.IsNull() {
		retriesStr := os.Getenv("CC_RETRIES")
		if retriesStr == "" {
			retries = 3
		} else {
			retries, _ = strconv.ParseInt(retriesStr, 0, 64)
		}
	} else {
		retries = config.Retries.ValueInt64()
	}

	var maxTimeout int64
	if config.MaxTimeout.IsUnknown() {
		// Cannot connect to client with an unknown value
		resp.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as max_timeout",
		)
		return
	}

	if config.MaxTimeout.IsNull() {
		maxTimeoutStr := os.Getenv("CC_MAX_TIMEOUT")
		if maxTimeoutStr == "" {
			maxTimeout = 60
		} else {
			maxTimeout, _ = strconv.ParseInt(maxTimeoutStr, 0, 64)
		}
	} else {
		maxTimeout = config.MaxTimeout.ValueInt64()
	}

	// Create a new catalyst center client and set it to the provider client
	c, err := cc.NewClient(url, username, password, cc.Insecure(insecure), cc.MaxRetries(int(retries)), cc.DefaultMaxAsyncWaitTime(int(maxTimeout)))
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create catalyst center client:\n\n"+err.Error(),
		)
		return
	}

	data := CcProviderData{Client: &c}
	resp.DataSourceData = &data
	resp.ResourceData = &data
}

func (p *CcProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAAASettingsResource,
		NewAnycastGatewayResource,
		NewAreaResource,
		NewAssignCredentialsResource,
		NewAssignTemplatesToTagResource,
		NewAssociateSiteToNetworkProfileResource,
		NewAuthenticationPolicyServerResource,
		NewBannerSettingsResource,
		NewBuildingResource,
		NewCredentialsCLIResource,
		NewCredentialsHTTPSReadResource,
		NewCredentialsHTTPSWriteResource,
		NewCredentialsSNMPv2ReadResource,
		NewCredentialsSNMPv2WriteResource,
		NewCredentialsSNMPv3Resource,
		NewDeployTemplateResource,
		NewDeviceResource,
		NewDeviceRoleResource,
		NewDHCPSettingsResource,
		NewDiscoveryResource,
		NewDNSSettingsResource,
		NewFabricAuthenticationProfileResource,
		NewFabricDeviceResource,
		NewFabricL2HandoffResource,
		NewFabricL2VirtualNetworkResource,
		NewFabricL3HandoffIPTransitResource,
		NewFabricL3VirtualNetworkResource,
		NewFabricPortAssignmentResource,
		NewFabricProvisionDeviceResource,
		NewFabricSiteResource,
		NewFabricVirtualNetworkResource,
		NewFabricVLANToSSIDResource,
		NewFloorResource,
		NewImageResource,
		NewImageActivationResource,
		NewImageDistributionResource,
		NewIPPoolResource,
		NewIPPoolReservationResource,
		NewLANAutomationResource,
		NewNetworkResource,
		NewNetworkProfileResource,
		NewNTPSettingsResource,
		NewPNPConfigPreviewResource,
		NewPnPDeviceResource,
		NewPnPDeviceClaimSiteResource,
		NewPnPImportDevicesResource,
		NewProjectResource,
		NewRoleResource,
		NewSPProfileResource,
		NewTagResource,
		NewTelemetrySettingsResource,
		NewTemplateResource,
		NewTemplateVersionResource,
		NewTimeZoneSettingsResource,
		NewTransitNetworkResource,
		NewUserResource,
		NewVirtualNetworkIPPoolResource,
		NewVirtualNetworkToFabricSiteResource,
		NewWirelessDeviceProvisionResource,
		NewWirelessEnterpriseSSIDResource,
		NewWirelessProfileResource,
		NewWirelessRFProfileResource,
		NewWirelessSSIDResource,
	}
}

func (p *CcProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewNetworkDevicesDataSource, // manually maintained
		NewAAASettingsDataSource,
		NewAnycastGatewayDataSource,
		NewAreaDataSource,
		NewAssignCredentialsDataSource,
		NewAssignTemplatesToTagDataSource,
		NewAuthenticationPolicyServerDataSource,
		NewBannerSettingsDataSource,
		NewBuildingDataSource,
		NewCredentialsCLIDataSource,
		NewCredentialsHTTPSReadDataSource,
		NewCredentialsHTTPSWriteDataSource,
		NewCredentialsSNMPv2ReadDataSource,
		NewCredentialsSNMPv2WriteDataSource,
		NewCredentialsSNMPv3DataSource,
		NewDeviceDetailDataSource,
		NewDHCPSettingsDataSource,
		NewDiscoveryDataSource,
		NewDNSSettingsDataSource,
		NewFabricAuthenticationProfileDataSource,
		NewFabricDeviceDataSource,
		NewFabricL2HandoffDataSource,
		NewFabricL2VirtualNetworkDataSource,
		NewFabricL3HandoffIPTransitDataSource,
		NewFabricL3VirtualNetworkDataSource,
		NewFabricPortAssignmentDataSource,
		NewFabricProvisionDeviceDataSource,
		NewFabricSiteDataSource,
		NewFabricVirtualNetworkDataSource,
		NewFabricVLANToSSIDDataSource,
		NewFloorDataSource,
		NewIPPoolDataSource,
		NewIPPoolReservationDataSource,
		NewLANAutomationDataSource,
		NewNetworkDataSource,
		NewNetworkProfileDataSource,
		NewNTPSettingsDataSource,
		NewPnPDeviceDataSource,
		NewProjectDataSource,
		NewRoleDataSource,
		NewSiteDataSource,
		NewSPProfileDataSource,
		NewTagDataSource,
		NewTelemetrySettingsDataSource,
		NewTemplateDataSource,
		NewTimeZoneSettingsDataSource,
		NewTransitNetworkDataSource,
		NewUserDataSource,
		NewWirelessEnterpriseSSIDDataSource,
		NewWirelessProfileDataSource,
		NewWirelessRFProfileDataSource,
		NewWirelessSSIDDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &CcProvider{
			version: version,
		}
	}
}

// End of section. //template:end provider
