// Package sandbox configures CoreWeave sandbox resources.
package sandbox

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the sandbox resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("coreweave_sandbox_managed_runner", func(r *config.Resource) {
		r.ShortGroup = "sandbox"
		r.Kind = "ManagedRunner"
		r.References["cluster_id"] = config.Reference{
			TerraformName: "coreweave_cks_cluster",
		}
	})
}
