package config

import (
	"fmt"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/pkg/errors"
)

// placeholderUUID is handed to the Terraform provider as the "id" of a
// resource whose identifier is assigned by the CoreWeave API, before the
// resource exists. The CoreWeave API answers NotFound for it, which the
// Terraform resource turns into an empty state, which Upjet reads as
// "does not exist yet, create it".
const placeholderUUID = "00000000-0000-0000-0000-000000000000"

// ExternalNameConfigs contains all external name configurations for this
// provider. Only resources listed here are generated.
//
// All CoreWeave resources are Terraform Plugin Framework resources, so they
// are reconciled in-process (no Terraform CLI). See provider.go.
var ExternalNameConfigs = map[string]config.ExternalName{
	// --- IDs assigned by the CoreWeave API (UUIDs) --------------------------
	// The external name is the API-assigned id; `name` stays a normal spec
	// field because it is not what the resource is looked up by.
	"coreweave_cks_cluster":                     computedID(),
	"coreweave_networking_vpc":                  computedID(),
	"coreweave_inference_capacity_claim":        computedID(),
	"coreweave_inference_deployment":            computedID(),
	"coreweave_inference_gateway":               computedID(),
	"coreweave_workload_federation_oidc_config": computedID(),
	// Access key IDs are assigned by the API too; the secret key is only
	// returned at creation and ends up in the connection secret.
	"coreweave_object_storage_access_key": computedID(),

	// --- Identified by a user-chosen parameter -------------------------------
	// Terraform id == runner_id, so the external name (defaulting to
	// metadata.name) is written into both.
	"coreweave_sandbox_managed_runner": config.ParameterAsIdentifier("runner_id"),

	// These resources have no "id" attribute at all; the identifying
	// attribute is read back from the Terraform state instead.
	"coreweave_object_storage_bucket":                     parameterFromState("name"),
	"coreweave_object_storage_organization_access_policy": parameterFromState("name"),

	// --- Bucket sub-resources --------------------------------------------------
	// Keyed by the bucket they configure. `bucket` stays in spec.forProvider so
	// it can be set with bucketRef/bucketSelector; the external name mirrors it.
	"coreweave_object_storage_bucket_policy":                  attributeFromState("bucket"),
	"coreweave_object_storage_bucket_settings":                attributeFromState("bucket"),
	"coreweave_object_storage_bucket_versioning":              attributeFromState("bucket"),
	"coreweave_object_storage_bucket_lifecycle_configuration": attributeFromState("bucket"),
	// Import ID is "<bucket>:<name>"; both stay in spec, the external name is
	// the inventory configuration name.
	"coreweave_object_storage_bucket_inventory": attributeFromState("name", "bucket"),
}

// computedID is for resources whose Terraform "id" is computed by the API.
func computedID() config.ExternalName {
	return config.FrameworkResourceWithComputedIdentifier("id", placeholderUUID)
}

// parameterFromState is config.ParameterAsIdentifier for resources without an
// "id" attribute: the external name is set into, and read back from, param.
func parameterFromState(param string) config.ExternalName {
	e := config.ParameterAsIdentifier(param)
	e.GetExternalNameFn = stateAttribute(param)
	return e
}

// attributeFromState keeps attr (and any extra identifier fields) in the spec
// and derives the external name from attr in the observed Terraform state.
func attributeFromState(attr string, extraIdentifierFields ...string) config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = stateAttribute(attr)
	e.IdentifierFields = append([]string{attr}, extraIdentifierFields...)
	return e
}

func stateAttribute(attr string) config.GetExternalNameFn {
	return func(tfstate map[string]any) (string, error) {
		if v, ok := tfstate[attr]; ok && v != nil {
			if s := fmt.Sprintf("%v", v); s != "" {
				return s, nil
			}
		}
		return "", errors.Errorf("cannot find attribute %q in tfstate", attr)
	}
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, 0, len(ExternalNameConfigs))
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l = append(l, name+"$")
	}
	return l
}
