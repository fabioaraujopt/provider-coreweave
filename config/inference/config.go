// Package inference configures CoreWeave Managed Inference resources.
package inference

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the inference resources.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("coreweave_inference_capacity_claim", func(r *config.Resource) {
		r.ShortGroup = "inference"
		r.Kind = "CapacityClaim"
	})
	p.AddResourceConfigurator("coreweave_inference_gateway", func(r *config.Resource) {
		r.ShortGroup = "inference"
		r.Kind = "Gateway"
	})
	p.AddResourceConfigurator("coreweave_inference_deployment", func(r *config.Resource) {
		r.ShortGroup = "inference"
		r.Kind = "Deployment"
		r.References["gateway_ids"] = config.Reference{
			TerraformName: "coreweave_inference_gateway",
		}
		r.References["model.bucket"] = config.Reference{
			TerraformName: "coreweave_object_storage_bucket",
		}
	})
}
