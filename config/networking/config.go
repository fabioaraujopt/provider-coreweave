// Package networking configures CoreWeave networking resources.
package networking

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the networking resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("coreweave_networking_vpc", func(r *config.Resource) {
		r.ShortGroup = "networking"
		r.Kind = "VPC"
	})
}
