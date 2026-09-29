// Package workloadfederation configures CoreWeave workload identity
// federation resources.
package workloadfederation

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the workload federation resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("coreweave_workload_federation_oidc_config", func(r *config.Resource) {
		r.ShortGroup = "workloadfederation"
		r.Kind = "OIDCConfig"
	})
}
