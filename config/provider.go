package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/coreweave/terraform-provider-coreweave/xpprovider"

	"github.com/your-org/provider-coreweave/config/cks"
	"github.com/your-org/provider-coreweave/config/inference"
	"github.com/your-org/provider-coreweave/config/networking"
	"github.com/your-org/provider-coreweave/config/objectstorage"
	"github.com/your-org/provider-coreweave/config/sandbox"
	"github.com/your-org/provider-coreweave/config/workloadfederation"
)

const (
	resourcePrefix = "coreweave"
	modulePath     = "github.com/your-org/provider-coreweave"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// configurators are shared by the cluster-scoped and namespaced providers.
// References use TerraformName, so the same configuration resolves to the
// right API package in both scopes.
var configurators = []func(provider *ujconfig.Provider){
	cks.Configure,
	networking.Configure,
	objectstorage.Configure,
	inference.Configure,
	sandbox.Configure,
	workloadfederation.Configure,
}

// GetProvider returns the cluster-scoped provider configuration
// (*.coreweave.crossplane.io, Crossplane v1-style managed resources).
func GetProvider() *ujconfig.Provider {
	return newProvider("coreweave.crossplane.io")
}

// GetProviderNamespaced returns the namespaced provider configuration
// (*.coreweave.m.crossplane.io, Crossplane v2 managed resources).
func GetProviderNamespaced() *ujconfig.Provider {
	return newProvider("coreweave.m.crossplane.io",
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))
}

func newProvider(rootGroup string, extra ...ujconfig.ProviderOption) *ujconfig.Provider {
	opts := append([]ujconfig.ProviderOption{
		ujconfig.WithRootGroup(rootGroup),
		ujconfig.WithShortName("coreweave"),
		// Every CoreWeave resource is built on the Terraform Plugin Framework,
		// so all of them are reconciled in-process through the framework
		// provider (Upjet "no-fork" architecture): no Terraform CLI, no
		// provider binary, no workspace on disk.
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginFrameworkIncludeList(ExternalNameConfigured()),
		ujconfig.WithTerraformPluginFrameworkProvider(xpprovider.New("")),
		ujconfig.WithFeaturesPackage("internal/features"),
		// Render Terraform blocks with MaxItems=1 as embedded objects rather
		// than single-element lists (the current Upjet default for new
		// providers; avoids a breaking API change later).
		ujconfig.WithSchemaTraversers(&ujconfig.SingletonListEmbedder{}),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
	}, extra...)

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata), opts...)
	for _, configure := range configurators {
		configure(pc)
	}
	pc.ConfigureResources()
	return pc
}
