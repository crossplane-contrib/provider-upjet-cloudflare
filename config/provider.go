package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	// The Terraform Plugin Framework provider, hosted in-process for the
	// resources listed in PluginFrameworkResources. See plugin_framework.go.
	cloudflare "github.com/cloudflare/terraform-provider-cloudflare/provider"

	// Cluster-scoped resource configurations
	accessCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/access"
	accountCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/account"
	apishieldCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/apishield"
	argoCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/argo"
	certificateCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/certificate"
	dnsCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/dns"
	emailCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/email"
	firewallCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/firewall"
	healthcheckCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/healthcheck"
	listCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/list"
	loadbalancerCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/loadbalancer"
	logpushCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/logpush"
	miscCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/misc"
	notificationCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/notification"
	r2Cluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/r2"
	spectrumCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/spectrum"
	streamCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/stream"
	waitingroomCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/waitingroom"
	workersCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/workers"
	zerotrustCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/zerotrust"
	zoneCluster "github.com/crossplane-contrib/provider-upjet-cloudflare/config/cluster/zone"

	// Namespaced resource configurations
	accessNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/access"
	accountNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/account"
	apishieldNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/apishield"
	argoNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/argo"
	certificateNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/certificate"
	dnsNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/dns"
	emailNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/email"
	firewallNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/firewall"
	healthcheckNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/healthcheck"
	listNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/list"
	loadbalancerNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/loadbalancer"
	logpushNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/logpush"
	miscNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/misc"
	notificationNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/notification"
	r2Namespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/r2"
	spectrumNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/spectrum"
	streamNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/stream"
	waitingroomNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/waitingroom"
	workersNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/workers"
	zerotrustNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/zerotrust"
	zoneNamespaced "github.com/crossplane-contrib/provider-upjet-cloudflare/config/namespaced/zone"
)

const (
	resourcePrefix = "cloudflare"
	modulePath     = "github.com/crossplane-contrib/provider-upjet-cloudflare"

	// providerVersion is reported to the hosted Terraform provider as its own
	// version. It is only surfaced in that provider's user agent, so it does
	// not need to track TERRAFORM_PROVIDER_VERSION.
	providerVersion = "crossplane"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("cloudflare.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithTerraformPluginFrameworkIncludeList(PluginFrameworkIncludeList()),
		ujconfig.WithTerraformPluginFrameworkProvider(cloudflare.NewProvider(providerVersion)),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		accountCluster.Configure,
		accessCluster.Configure,
		apishieldCluster.Configure,
		argoCluster.Configure,
		certificateCluster.Configure,
		dnsCluster.Configure,
		emailCluster.Configure,
		firewallCluster.Configure,
		healthcheckCluster.Configure,
		listCluster.Configure,
		loadbalancerCluster.Configure,
		logpushCluster.Configure,
		miscCluster.Configure,
		notificationCluster.Configure,
		r2Cluster.Configure,
		spectrumCluster.Configure,
		streamCluster.Configure,
		waitingroomCluster.Configure,
		workersCluster.Configure,
		zerotrustCluster.Configure,
		zoneCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("cloudflare.m.cloudflare.com"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithTerraformPluginFrameworkIncludeList(PluginFrameworkIncludeList()),
		ujconfig.WithTerraformPluginFrameworkProvider(cloudflare.NewProvider(providerVersion)),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		accountNamespaced.Configure,
		accessNamespaced.Configure,
		apishieldNamespaced.Configure,
		argoNamespaced.Configure,
		certificateNamespaced.Configure,
		dnsNamespaced.Configure,
		emailNamespaced.Configure,
		firewallNamespaced.Configure,
		healthcheckNamespaced.Configure,
		listNamespaced.Configure,
		loadbalancerNamespaced.Configure,
		logpushNamespaced.Configure,
		miscNamespaced.Configure,
		notificationNamespaced.Configure,
		r2Namespaced.Configure,
		spectrumNamespaced.Configure,
		streamNamespaced.Configure,
		waitingroomNamespaced.Configure,
		workersNamespaced.Configure,
		zerotrustNamespaced.Configure,
		zoneNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
