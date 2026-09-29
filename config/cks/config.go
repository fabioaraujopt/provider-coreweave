// Package cks configures CoreWeave Kubernetes Service (CKS) resources.
package cks

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the CKS resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("coreweave_cks_cluster", func(r *config.Resource) {
		r.ShortGroup = "cks"
		r.Kind = "Cluster"
		r.References["vpc_id"] = config.Reference{
			TerraformName: "coreweave_networking_vpc",
		}
		// shared_storage_cluster_id points at another CKS cluster.
		r.References["shared_storage_cluster_id"] = config.Reference{
			TerraformName:     "coreweave_cks_cluster",
			RefFieldName:      "SharedStorageClusterRef",
			SelectorFieldName: "SharedStorageClusterSelector",
		}
	})
}
