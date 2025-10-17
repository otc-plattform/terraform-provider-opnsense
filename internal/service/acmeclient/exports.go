package acmeclient

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		newACMEClientAccountResource,
		newACMEClientChallengeResource,
		newACMEClientAutomationResource,
		newACMEClientCertificateResource,
		newACMEClientSettingsResource,
	}
}

func DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		newACMEClientAccountDataSource,
		newACMEClientChallengeDataSource,
		newACMEClientAutomationDataSource,
		newACMEClientCertificateDataSource,
		newACMEClientSettingsDataSource,
	}
}
